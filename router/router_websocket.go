package router

import (
	"github.com/gin-gonic/gin"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/v1/api"
)

func InitWebSocketRouter(Router *gin.RouterGroup) {
	newServer := wsmanage.NewServer()
	UserRouter := Router.Group("")
	{

		global.GVA_LOG.Info("register wsmanage handler")

		//配置路由
		newServer.AddRouter("ping", v1.PingRouter{})        //ping保持连接
		newServer.AddRouter("login", v1.Login{})            //登录
		newServer.AddRouter("add_friend", v1.AddFriend{})   //添加好友
		newServer.AddRouter("list_friend", v1.ListFriend{}) //好友列表

		newServer.AddRouter("create_group", v1.CreateGroup{}) //创建群
		newServer.AddRouter("list_group", v1.ListGroup{})     //群列表
		newServer.AddRouter("join_group", v1.JoinGroup{})     //加群

		//即时通讯文字api
		newServer.AddRouter("single_msg", v1.SingleMsg{}) //单聊
		newServer.AddRouter("room_msg", v1.RoomMsg{})     //群聊

		//即时通讯音视频api p2p视频

		newServer.AddRouter("offer", v1.SingleCandidate{}) //群聊
		newServer.AddRouter("answer", v1.SingleCandidate{})
		newServer.AddRouter("candidate", v1.SingleCandidate{})

		UserRouter.GET("/ws", newServer.Serve) //启动连接

	}
}
