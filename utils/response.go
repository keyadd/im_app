package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"im_app/global"
	"im_app/initialize/wsmanage"
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

func ResponseSuccess(r wsmanage.Request, c *wsmanage.Connection, data interface{}) {
	rd := &ResponseData{
		Code: global.CodeSuccess,
		Msg:  global.CodeSuccess.Msg(),
		Data: data,
	}
	err := c.SendMessage(r, rd)
	if err != nil {
		global.GVA_LOG.Error("err = c.SendMessage(r, marshal)", zap.Error(err))
	}
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
