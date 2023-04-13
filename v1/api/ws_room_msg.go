package v1

import (
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/middleware"
	"im_app/utils"
	"im_app/v1/model/request"
	"im_app/v1/service"
)

type RoomMsg struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (t RoomMsg) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (t RoomMsg) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()
	_, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	msg := request.RoomMessage{}
	res, err := utils.MapToStruct(data, &msg)
	if err != nil {
		global.GVA_LOG.Error("utils.MapToStruct(data, &joinDate)", zap.Error(err))
		return
	}
	message := res.(*request.RoomMessage)

	ids := service.GetUserArr(message.GroupNo)

	//用户id 连接信息添加
	mgr := c.Server.GetConnMgr()
	for _, v := range ids {

		conn, err := mgr.GetConn(v)

		if err != nil {
			global.GVA_LOG.Error("conn, err := mgr.GetConn(msg.FriendId)", zap.Error(err))
			return
		}
		utils.ResponseSuccess(r, conn, message.Content)
	}

}
