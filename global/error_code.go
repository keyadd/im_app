package global

import "errors"

var (
	ErrorUserExist        = errors.New("用户已存在")
	ErrorUserNotExist     = errors.New("not is username")
	ErrorPasswordNotExist = errors.New("not is password")
	ErrorInvalidID        = errors.New("无效的ID")
	//Error
)
