package service

import (
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize"
	"im_app/v1/model"
	"im_app/v1/model/request"
	"im_app/v1/model/response"
	"time"
)

func CreateGroup(id int64, r *request.GroupRequest) error {

	uuid := initialize.GenID()

	group := response.Group{
		UserId:     id,
		Name:       r.Name,
		GroupNo:    uuid,
		CreateTime: time.Now().Unix(),
	}
	g, err := model.CreateGroup(group)

	if err != nil {
		global.GVA_LOG.Error("model.CreateGroup(group)", zap.Error(err))
		return err
	}
	groupRequest := request.UserGroupRequest{
		UserId:     id,
		GroupId:    g.Id,
		CreateTime: time.Now().Unix(),
	}
	err = model.AddUserGroup(groupRequest)
	if err != nil {
		global.GVA_LOG.Error("count, err := model.GetGroupNo(r)", zap.Error(err))
		return err
	}
	return nil

}
