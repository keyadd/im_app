package wsmanage

type Message struct {
	MsgID string `json:"msg_id"` //业务消息ID
	Data  string `json:"data"`   //消息的内容
}

//创建一个Message消息包
func NewMsg(msgID string, data string) *Message {
	return &Message{
		MsgID: msgID,
		Data:  data,
	}
}

////获取消息websocket类型
//func (msg *Message) GetMsgType() int {
//	return msg.MsgType.MsgType
//}

//获取消息
func (msg *Message) GetMsgID() string {
	return msg.MsgID
}

//获取消息内容
func (msg *Message) GetData() string {
	return msg.Data
}
