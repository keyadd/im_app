package v1

import (
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/middleware"
	"im_app/utils"
	"im_app/v1/model/request"
)

type SingleCandidate struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (s SingleCandidate) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (s SingleCandidate) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()
	_, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	msg := request.Message{}
	res, err := utils.MapToStruct(data, &msg)
	if err != nil {
		global.GVA_LOG.Error("utils.MapToStruct(data, &joinDate)", zap.Error(err))
		return
	}
	message := res.(*request.Message)

	//用户id 连接信息添加
	mgr := c.Server.GetConnMgr()
	conn, err := mgr.GetConn(message.FriendId)

	if err != nil {
		global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
		return
	}
	utils.ResponseSuccess(r, conn, message.Content)

}
