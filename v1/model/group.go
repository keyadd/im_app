package model

import (
	"im_app/global"
	"im_app/v1/model/response"
)

func CreateGroup(g response.Group) (response.Group, error) {
	db := global.GVA_DB.Table("group")
	err := db.Create(&g).Error
	return g, err
}

func ListGroup(user_id int64) (l []response.GroupList, err error) {
	db := global.GVA_DB.Table("group")
	err = db.Where("user_id = ?", user_id).Find(&l).Error

	return l, err
}

func GetGroupNo(group_no int64) (l response.GroupList, err error) {
	db := global.GVA_DB.Table("group")
	err = db.Where("group_no = ?", group_no).Find(&l).Error
	return l, err
}
