package v1

import (
	"im_app/initialize/wsmanage"
)

type PingRouter struct {
	*wsmanage.BaseRouter
}

// Ping Handle
func (p PingRouter) Handle(request wsmanage.Request) {
	//data := request.GetData()
	//
	//if data == "ping" {
	//	connection := request.GetConnection()
	//	err := connection.SendMessage(request, "pong")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}

}
