package wsmanage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"im_app/global"
	"net/http"
	"sync/atomic"
	"time"
)

type Server struct {
	//Server的消息管理模块
	MsgHandler MsgHandle
	//当前Server链接管理器
	ConnMgr ConnManager
	//当前Server连接创建时的hook函数
	OnConnStart func(conn Connection)
	//当前Server连接断开时的hook函数
	OnConnStop func(conn Connection)
}

func NewServer() *Server {
	return &Server{
		ConnMgr:    NewConnManager(),
		MsgHandler: NewMsgHandle(),
	}
}
func (s *Server) Start(w http.ResponseWriter, r *http.Request) {
	//开启一个go去做服务端Linster业务
	go func() {

		//TODO wsmanage.go 应该有一个自动生成ID的方法
		curConnId := uint64(time.Now().Unix())
		connId := atomic.AddUint64(&curConnId, 1)

		//初始化连接

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(s.ConnMgr.Len())

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.ConnMgr.Len() >= global.GVA_CONFIG.WebSocket.MaxConnLen {
			conn.Close()
			return
		}

		dealConn := NewConnection(*s, conn, connId, s.MsgHandler, r)
		fmt.Println("Current connId:", connId)
		//3.4 启动当前链接的处理业务
		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn.Start()

	}()
}

// 停止服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Websocket wsmanage , name ")

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

// 运行服务
func (s *Server) Serve(c *gin.Context) {
	s.Start(c.Writer, c.Request)

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	select {}
}

// 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(msgId string, router IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}

// 得到链接管理
func (s *Server) GetConnMgr() ConnManager {
	return s.ConnMgr
}

// 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(Connection)) {
	s.OnConnStart = hookFunc
}

// 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(Connection)) {
	s.OnConnStop = hookFunc
}

// 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn Connection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn Connection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
