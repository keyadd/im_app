package model

import (
	"im_app/global"
	"im_app/v1/model/request"
	"im_app/v1/model/response"
)

func GetUser(user request.UserLogin) (m response.User, err error) {
	db := global.GVA_DB.Table("user")
	u := response.User{}
	err = db.Where("email = ?", user.Email).First(&u).Error
	return u, err
}

func GetUserCount(user_id int64, friend_id int64) (count int64) {
	db := global.GVA_DB.Table("user")
	db.Where("user_id = ?", user_id).Where("user_friend_id = ?", friend_id).Count(&count)
	return
}
