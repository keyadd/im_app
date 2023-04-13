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

type JoinGroup struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (a JoinGroup) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (a JoinGroup) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()
	user_id, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	joinDate := request.JoinGroup{}
	res, err := utils.MapToStruct(data, &joinDate)
	if err != nil {
		global.GVA_LOG.Error("utils.MapToStruct(data, &joinDate)", zap.Error(err))
		return
	}
	joinGroupRes := res.(*request.JoinGroup)

	err = service.JoinGroup(user_id, joinGroupRes)
	if err != nil {
		global.GVA_LOG.Error("err = service.AddFriend(user_id, friend.FriendId)", zap.Error(err))
		utils.ResponseSuccess(r, &c, err.Error())
		return
	}
	utils.ResponseSuccess(r, &c, "添加成功")
}
