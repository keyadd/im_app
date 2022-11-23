package wsmanage

import (
	"app_ws/global"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/websocket"
	"net"
	"net/http"
	"sync"
)

type Connection struct {
	//当前链接所属Server
	Server      Server
	Conn        *websocket.Conn
	HttpRequest *http.Request
	connId      uint64
	//inChan            chan *Message
	outChan  chan *Message
	isClosed bool
	//告知当前连接已经退出/停止 channel
	ExitChan  chan bool
	MsgHandle MsgHandle
	//链接属性
	property map[string]interface{}
	//保护链接属性修改的锁
	propertyLock sync.RWMutex

	mutex sync.Mutex
}

//初始化链接服务
func NewConnection(server Server, wsSocket *websocket.Conn, connId uint64, msgHandler MsgHandle, r *http.Request) *Connection {
	c := Connection{
		Server:      server,
		Conn:        wsSocket,
		connId:      connId,
		MsgHandle:   msgHandler,
		isClosed:    false,
		outChan:     make(chan *Message, 1024),
		ExitChan:    make(chan bool, 1),
		HttpRequest: r,
	}
	mgr := c.Server.GetConnMgr()
	mgr.Add(c)
	return &c
}

//开始
func (c *Connection) Start() {

	go c.readLoop()
	go c.writeLoop()
	c.Server.CallOnConnStart(*c)
	select {}
}

//关闭连接
func (c *Connection) Close() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.isClosed {
		c.isClosed = true
		c.Server.CallOnConnStop(*c)
		close(c.ExitChan)
	}

	c.Conn.Close()

	mgr := c.Server.GetConnMgr()
	mgr.Remove(*c)
}
func (c *Connection) GetConnectionState() bool {
	return c.isClosed
}

func (c *Connection) GetHttpRequest() *http.Request {
	return c.HttpRequest
}

//获取链接对象
func (c *Connection) GetConnection() *websocket.Conn {
	return c.Conn
}

//获取链接ID
func (c *Connection) GetConnID() uint64 {
	return c.connId
}

//获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//设置链接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

//获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

//移除链接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

//读websocket
func (c *Connection) readLoop() {

	for {
		msg := Message{}
		var reply []byte
		err := websocket.Message.Receive(c.Conn, &reply)
		if err != nil {
			goto ERR
		}
		if err := json.Unmarshal(reply, &msg); err != nil {
			fmt.Println("Error:", err)
			if err != nil {
				goto ERR
			}
		}
		if msg.MsgID != "" {
			//fmt.Println(msg.MsgID)
			message := NewMsg(msg.MsgID, msg.Data)
			//得到当前客户端请求的Request数据
			req := Request{
				conn: *c,
				msg:  *message,
			}
			if global.GVA_CONFIG.WebSocket.WorkerPoolSize > 1000 {
				//已经启动工作池机制，将消息交给Worker处理
				c.MsgHandle.SendMsgToTaskQueue(req)
			} else {
				//从绑定好的消息和对应的处理方法中执行对应的Handle方法
				go c.MsgHandle.DoMsgHandler(req)
			}

		} else {
			fmt.Println("消息标识MsgID不存在!")

		}

	}

ERR:
	c.Close()
}

//写websocket
func (c *Connection) writeLoop() {

	for {
		select {
		case message := <-c.outChan:
			msg, err := json.Marshal(message)
			//fmt.Println(string(msg))
			if err != nil {
				fmt.Println(err)
				goto ERR
			}
			err = websocket.Message.Send(c.Conn, string(msg))
			if err != nil {
				fmt.Println(err)
				goto ERR
			}
		case <-c.ExitChan:
			goto CLOSED
		}
	}
ERR:
	c.Close()
CLOSED:
}

// SendMessage 发送消息
func (c *Connection) SendMessage(request Request, msgData string) (err error) {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	message := NewMsg(request.GetMsgID(), msgData)

	c.outChan <- message
	return nil
}
