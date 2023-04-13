package v1

import (
	"fmt"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/middleware"
	"im_app/utils"
	"im_app/v1/model/request"
	"im_app/v1/service"
)

type CreateGroup struct {
	*wsmanage.BaseRouter
}

// PreHandle 前置拦截器 处理token
func (l CreateGroup) PreHandle(r wsmanage.Request) error {
	err := middleware.WsJWTMiddleware(r)
	return err
}

func (l CreateGroup) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()
	user_id, is := c.Get(global.GVA_CONFIG.JWT.UserIdName)
	if is != true {
		global.GVA_LOG.Error("获取不到map中user_id")
		utils.ResponseSuccess(r, &c, "token失效")
		return
	}
	fmt.Println(data)

	groupDate := request.GroupRequest{}
	res, err := utils.MapToStruct(data, &groupDate)
	if err != nil {
		global.GVA_LOG.Error("utils.MapToStruct(data, &joinDate)", zap.Error(err))
		return
	}
	groupRes := res.(*request.GroupRequest)

	err = service.CreateGroup(user_id.(int64), groupRes)
	if err != nil {
		utils.ResponseSuccess(r, &c, "创建失败")
	}

	utils.ResponseSuccess(r, &c, "创建成功")

}
