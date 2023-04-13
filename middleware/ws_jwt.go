package middleware

import (
	"errors"
	"im_app/global"
	"im_app/initialize"
	"im_app/initialize/wsmanage"
	"im_app/utils"
)

func WsJWTMiddleware(r wsmanage.Request) error {

	data := r.GetToken()
	//fmt.Println(data)
	c := r.GetConnection()
	if data == "" {
		global.GVA_LOG.Error("获取header 出错")
		return errors.New("获取header 出错")
	}
	mc, err := initialize.ParseToken(data)
	if err != nil {
		global.GVA_LOG.Error("token失效")
		utils.ResponseSuccess(r, &c, "token失效")
		return err
	}
	UserIdName := global.GVA_CONFIG.JWT.UserIdName
	//c := r.GetConnection()
	c.Set(UserIdName, mc.UserID)
	return nil
}
