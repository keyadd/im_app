package v1

import (
	"app_ws/initialize/wsmanage"
	"log"
)

type EchoRouter struct {
	*wsmanage.BaseRouter
}

func (h EchoRouter) Handle(request wsmanage.Request) {
	data := request.GetData()
	if data != "" {
		connection := request.GetConnection()
		err := connection.SendMessage(request, data)
		if err != nil {
			//zlog.Error(err)
			log.Fatalln(err)
		}
	}

}
