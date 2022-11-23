package v1

import (
	"app_ws/core/wsmanage"
	"log"
)

type PingRouter struct {
	*wsmanage.BaseRouter
}

//Ping Handle
func (p PingRouter) Handle(request wsmanage.Request) {
	data := request.GetData()
	//fmt.Println(data)

	if data == "ping" {
		connection := request.GetConnection()
		err := connection.SendMessage(request, "pong")
		if err != nil {
			log.Println(err)
		}
	}

}
