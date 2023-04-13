package wsmanage

type Request struct {
	conn        Connection     //已经和客户端建立好的 链接
	msg         RequestMessage //客户端请求的数据
	connManager ConnManager    //连接管理
}

// 获取请求连接信息
func (r *Request) GetConnection() Connection {
	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() any {
	return r.msg.GetData()
}
func (r *Request) GetToken() string {
	return r.msg.GetToken()
}

// 获取请求的消息的ID
func (r *Request) GetMsgID() string {
	return r.msg.GetMsgID()
}
