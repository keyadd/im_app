package model

import (
	"im_app/global"
	"im_app/v1/model/request"
	"im_app/v1/model/response"
)

func GetUserFriendCount(user_id int64, friend_id int64) (count int64) {
	db := global.GVA_DB.Table("friend")
	db.Where("user_id = ?", user_id).Where("user_friend_id = ?", friend_id).Count(&count)
	return
}

func CreateUserFriend(user_friend request.Friend, friend_user request.Friend) error {
	db := global.GVA_DB.Table("friend")
	tx := db.Begin()
	err := tx.Create(&user_friend).Error
	if err != nil {
		tx.Rollback()
	}
	err = tx.Create(&friend_user).Error
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error
}

func GetFriendList(id int64) (u []response.FriendList, err error) {
	//var u []response.FriendList
	db := global.GVA_DB.Table("friend")
	err = db.Where("user_id = ?", id).Preload("User").Find(&u).Error

	return u, err

}
