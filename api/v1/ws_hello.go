package v1

import (
	"app_ws/core/wsmanage"
	"errors"
	"fmt"
	"log"
	"time"
)

type HelloRouter struct {
	*wsmanage.BaseRouter
}

type UserInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (h HelloRouter) Handle(request wsmanage.Request) {
	for {
		now := time.Now()
		newtime := now.Format("2006-01-02 03:04:05")
		connection := request.GetConnection()
		err := connection.SendMessage(request, newtime)
		if err != nil {
			//zlog.Error(err)
			break
		}
		time.Sleep(time.Second * 1)

	}
}

func (h HelloRouter) PreHandle(request wsmanage.Request) error {
	connection := request.GetConnection()
	data := connection.GetHttpRequest()
	getUserAgent := data.Header.Get("User-Agent")
	fmt.Println(getUserAgent)
	err := connection.SendMessage(request, "消息错误")
	if err != nil {
		log.Println(err)
	}

	if getUserAgent == "" {
		return errors.New("获取header 出错")
	}
	return nil

}
