package wsmanage

type RequestMessage struct {
	MsgID string `json:"msg_id"` //业务消息ID
	Data  any    `json:"data"`   //消息的内容
	Token string `json:"token"`
}
type ResponseMessage struct {
	MsgID string `json:"msg_id"` //业务消息ID
	Data  any    `json:"data"`   //消息的内容
}

func NewRequestMessage(msgID string, data any, token string) *RequestMessage {
	return &RequestMessage{
		MsgID: msgID,
		Data:  data,
		Token: token,
	}
}

// 创建一个Message消息包
func NewResponseMessage(msgID string, data any) *ResponseMessage {
	return &ResponseMessage{
		MsgID: msgID,
		Data:  data,
	}
}

////获取消息websocket类型
//func (msg *Message) GetMsgType() int {
//	return msg.MsgType.MsgType
//}

// 获取消息
func (msg *RequestMessage) GetMsgID() string {
	return msg.MsgID
}
func (msg *RequestMessage) GetToken() string {
	return msg.Token
}

// 获取消息内容
func (msg *RequestMessage) GetData() any {
	return msg.Data
}
