package wsmanage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"im_app/global"
	"net"
	"net/http"
	"sync"
)

type Connection struct {
	//当前链接所属Server
	Server      Server
	Conn        net.Conn
	HttpRequest *http.Request
	connId      uint64
	//inChan            chan *Message
	outChan  chan *ResponseMessage
	isClosed bool
	//告知当前连接已经退出/停止 channel
	ExitChan  chan bool
	MsgHandle MsgHandle
	//存储当前连接的用户信息
	Keys map[string]any
	//保护链接属性修改的锁
	mu sync.RWMutex

	mutex sync.Mutex
}

// 初始化链接服务
func NewConnection(server Server, conn net.Conn, connId uint64, msgHandler MsgHandle, r *http.Request) *Connection {
	c := Connection{
		Server:      server,
		Conn:        conn,
		connId:      connId,
		MsgHandle:   msgHandler,
		isClosed:    false,
		outChan:     make(chan *ResponseMessage, 1024),
		ExitChan:    make(chan bool, 1),
		HttpRequest: r,
		Keys:        make(map[string]any),
	}
	mgr := c.Server.GetConnMgr()
	mgr.Add(c)
	return &c
}

// 开始
func (c *Connection) Start() {

	go c.readLoop()
	go c.writeLoop()
	c.Server.CallOnConnStart(*c)
	select {}
}

// 关闭连接
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

// 获取链接对象
func (c *Connection) GetConnection() net.Conn {
	return c.Conn
}

// 获取链接ID
func (c *Connection) GetConnID() uint64 {
	return c.connId
}

// 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 设置用户信息
func (c *Connection) Set(key string, value any) {
	c.mu.Lock()
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
	c.mu.Unlock()
}

// 获取用户信息
func (c *Connection) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	value, exists = c.Keys[key]
	c.mu.RUnlock()
	return
}

// 读websocket
func (c *Connection) readLoop() {
	r := wsutil.NewReader(c.Conn, ws.StateServerSide)
	decoder := json.NewDecoder(r)
	for {
		hdr, err := r.NextFrame()
		if err != nil {
			goto ERR
		}
		msg := RequestMessage{}
		if hdr.OpCode == ws.OpClose {
			goto ERR
		}

		if err = decoder.Decode(&msg); err != nil {
			//return err
			goto ERR
		}
		fmt.Println(msg)
		if msg.MsgID != "" {
			//fmt.Println(msg.MsgID)
			message := NewRequestMessage(msg.MsgID, msg.Data, msg.Token)
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

// 写websocket
func (c *Connection) writeLoop() {

	w := wsutil.NewWriter(c.Conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	for {
		select {
		case message := <-c.outChan:
			if err := encoder.Encode(&message); err != nil {
				goto ERR
			}
			if err := w.Flush(); err != nil {
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
func (c *Connection) SendMessage(request Request, msgData any) (err error) {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	message := NewResponseMessage(request.GetMsgID(), msgData)

	c.outChan <- message
	return nil
}
