package v1

import (
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/middleware"
	"im_app/utils"
	"im_app/v1/service"
)

type ListGroup struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (l ListGroup) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (l ListGroup) Handle(r wsmanage.Request) {
	c := r.GetConnection()
	user_id, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	list := service.GetGroupList(user_id)
	utils.ResponseSuccess(r, &c, list)

}
