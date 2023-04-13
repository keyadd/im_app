package model

import (
	"im_app/global"
	"im_app/v1/model/request"
)

func AddUserGroup(r request.UserGroupRequest) (err error) {
	db := global.GVA_DB.Table("user_group")
	err = db.Create(&r).Error
	return
}

func GetUserIdArray(group_id int64) (userIdArr []int64) {
	db := global.GVA_DB.Table("user_group")
	db.Where("group_id = ?", group_id).Select("user_id").Find(&userIdArr)
	return
}
