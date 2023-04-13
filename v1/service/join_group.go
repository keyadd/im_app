package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/v1/model"
	"im_app/v1/model/request"
	"time"
)

var (
	ErrorJoinGroup = errors.New("无法搜到群")
)

func JoinGroup(id any, r *request.JoinGroup) error {
	user_id := id.(int64)
	fmt.Println(r.GroupNo)
	l, err := model.GetGroupNo(r.GroupNo)
	if l.Id == 0 && err != nil {
		global.GVA_LOG.Error("count, err := model.GetGroupNo(r)", zap.Error(err))
		return err
	}

	groupRequest := request.UserGroupRequest{
		UserId:     user_id,
		GroupId:    l.Id,
		CreateTime: time.Now().Unix(),
	}
	err = model.AddUserGroup(groupRequest)
	if err != nil {
		global.GVA_LOG.Error("count, err := model.GetGroupNo(r)", zap.Error(err))
		return err
	}
	return nil

}
