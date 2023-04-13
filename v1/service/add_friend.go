package service

import (
	"errors"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/v1/model"
	"im_app/v1/model/request"
	"time"
)

var (
	ErrorUserIsFriend = errors.New("用户已经是好友")
	ErrorUserCreate   = errors.New("添加好友失败")
)

func AddFriend(user_id any, friend_id int64) error {

	id := user_id.(int64)
	count := model.GetUserFriendCount(id, friend_id)
	global.GVA_LOG.Error("count := model.GetUserFriendCount(id, friend_id)")
	if count > 0 {
		return ErrorUserIsFriend
	}
	times := time.Now().Unix()

	user_friend := request.Friend{UserId: id, UserFriendId: friend_id, CreateTime: times}
	friend_user := request.Friend{UserId: friend_id, UserFriendId: id, CreateTime: times}

	err := model.CreateUserFriend(user_friend, friend_user)
	if err != nil {
		global.GVA_LOG.Error("model.CreateUserFriend(user_friend, friend_user)", zap.Error(err))
		return ErrorUserCreate
	}
	return nil

}
