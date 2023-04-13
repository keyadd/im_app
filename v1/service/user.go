package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize"
	"im_app/utils"
	"im_app/v1/model"
	"im_app/v1/model/request"
)

var (
	//ErrorUserExist        = errors.New("用户已存在")
	ErrorUserNotExist     = errors.New("not is username")
	ErrorPasswordNotExist = errors.New("not is password")
)

func Login(r request.UserLogin) (token string, id int64, err error) {
	m, err := model.GetUser(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetUser(r)", zap.Error(err))
		return
	}

	check, err := utils.CheckPwd(r.Password, m.Password)
	fmt.Println(check)
	if err != nil {
		global.GVA_LOG.Error("utils.CheckPwd(r.Password, m.Password)", zap.Error(err))
		return
	}
	if check != true {
		return "", 0, ErrorPasswordNotExist
	}
	genToken, err := initialize.GenToken(m.Id, m.Email)
	if err != nil {
		global.GVA_LOG.Error("genToken, err := initialize.GenToken(m.Id, m.Email)", zap.Error(err))
		return
	}
	return genToken, m.Id, err

}
