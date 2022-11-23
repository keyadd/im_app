package utils

import (
	"app_ws/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ResponseData struct {
	Code global.ResCode `json:"code"`
	Msg  interface{}    `json:"msg"`
	Data interface{}    `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code global.ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: global.CodeSuccess,
		Msg:  global.CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code global.ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsgValid(c *gin.Context, code global.ResCode, msg interface{}) {

	rd := &ResponseData{
		Code: code,
		Msg:  RemoveTopStruct(msg.(validator.ValidationErrorsTranslations)),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
