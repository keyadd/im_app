package service

import (
	"go.uber.org/zap"
	"im_app/global"
	"im_app/v1/model"
	"im_app/v1/model/response"
)

func GetGroupList(id any) (list []response.GroupList) {
	user_id := id.(int64)
	u, err := model.ListGroup(user_id)
	if err != nil {
		global.GVA_LOG.Error("model.ListGroup(user_id)", zap.Error(err))
		return
	}
	return u

}
