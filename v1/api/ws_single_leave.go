package v1

import (
	"encoding/json"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/middleware"
	"im_app/utils"
	"im_app/v1/model/request"
)

type Leave struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (l Leave) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (l Leave) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()
	userId, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	msg := request.Publish{}
	res, err := utils.MapToStruct(data, &msg)
	if err != nil {
		global.GVA_LOG.Error("utils.MapToStruct(data, &joinDate)", zap.Error(err))
		return
	}
	message := res.(*request.Publish)

	c.AddWebRTCPeer(userId.(string), true)

	resp := make(map[string]interface{})
	resp["leave"] = userId
	resp["status"] = "leave"
	respByte, err := json.Marshal(resp)
	if err != nil {
		return
	}
	respStr := string(respByte)

	if respStr != "" {
		mgr := c.Server.GetConnMgr()

		iconn, err := mgr.GetConn(userId.(int64))
		if err != nil {
			global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
			return
		}
		utils.ResponseSuccess(r, iconn, respStr)

		conn, err := mgr.GetConn(message.FriendId)
		if err != nil {
			global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
			return
		}

		utils.ResponseSuccess(r, conn, respStr)

		c.DelWebRTCPeer(userId.(string), true)
		c.DelWebRTCPeer(userId.(string), false)

	}

}
