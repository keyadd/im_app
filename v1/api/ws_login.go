package v1

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
	"im_app/utils"
	"im_app/v1/model/request"
	"im_app/v1/service"
)

type Login struct {
	*wsmanage.BaseRouter
}

func (h Login) Handle(r wsmanage.Request) {
	data := r.GetData()
	c := r.GetConnection()

	login := request.UserLogin{}
	fmt.Println(data)
	err := mapstructure.Decode(data, &login)
	if err != nil {
		global.GVA_LOG.Error("mapstructure.Decode(data, &login)", zap.Error(err))
	}
	token, id, err := service.Login(login)
	if err != nil {
		utils.ResponseSuccess(r, &c, err)
		return
	}

	//用户id 连接信息添加
	mgr := c.Server.GetConnMgr()
	mgr.AddConn(id, r.GetConnection())
	utils.ResponseSuccess(r, &c, token)

}
