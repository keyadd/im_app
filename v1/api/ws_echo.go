package v1

import (
	"im_app/initialize/wsmanage"
)

type EchoRouter struct {
	*wsmanage.BaseRouter
}
type Echo struct {
	data string
}

func (h EchoRouter) Handle(request wsmanage.Request) {
	//data := request.GetData()
	//echo := Echo{}
	//json.Unmarshal(data, echo)
	//
	////fmt.Println(value)
	//if data != "" {
	//	connection := request.GetConnection()
	//	value, _ := connection.Get("User-Agent")
	//	fmt.Println(value)
	//	err := connection.SendMessage(request, data)
	//	if err != nil {
	//		//zlog.Error(err)
	//		log.Fatalln(err)
	//	}
	//}

}

func (h EchoRouter) PreHandle(request wsmanage.Request) error {
	//connection := request.GetConnection()
	//
	//data := connection.GetHttpRequest()
	//getUserAgent := data.Header.Get("User-Agent")
	//connection.Set("User-Agent", getUserAgent)
	////fmt.Println(getUserAgent)
	//fmt.Println("---------------------------------------------------------------------------------------")
	//err := connection.SendMessage(request, "消息错误")
	//if err != nil {
	//	log.Println(err)
	//}
	////h.Set("User-Agent", getUserAgent)
	//
	//if getUserAgent == "" {
	//	return errors.New("获取header 出错")
	//}
	return nil

}
