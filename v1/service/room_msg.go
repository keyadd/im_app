package service

import (
	"go.uber.org/zap"
	"im_app/global"
	"im_app/v1/model"
)

func GetUserArr(group_no int64) (ids []int64) {
	l, err := model.GetGroupNo(group_no)
	if err != nil {
		global.GVA_LOG.Error("count, err := model.GetGroupNo(r)", zap.Error(err))
		return
	}

	arr := model.GetUserIdArray(l.Id)

	return arr

}
