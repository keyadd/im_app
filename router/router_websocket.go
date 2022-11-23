package router

import (
	v1 "app_ws/api/v1"
	"app_ws/core/wsmanage"
	"app_ws/global"
	"github.com/gin-gonic/gin"
)

func InitWebSocketRouter(Router *gin.RouterGroup) {
	newServer := wsmanage.NewServer()
	UserRouter := Router.Group("")
	{

		global.GVA_LOG.Info("register wsmanage handler")

		//配置路由
		newServer.AddRouter("ping", v1.PingRouter{}) //ping保持连接
		//newServer.SetOnConnStart(DoConnectionBegin)
		//newServer.SetOnConnStop(DoConnectionEnd)
		newServer.AddRouter("hello", v1.HelloRouter{})
		newServer.AddRouter("echo", v1.EchoRouter{})
		UserRouter.GET("/ws", newServer.Serve)
	}
}
