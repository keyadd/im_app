package response

type Group struct {
	Id         int64  `json:"id"`
	GroupNo    int64  `json:"group_no"`
	Name       string `json:"name"`
	UserId     int64  `json:"user_id"`
	CreateTime int64  `json:"create_time"`
}

type GroupList struct {
	Id      int64  `json:"id"`
	GroupNo int64  `json:"group_no"`
	Name    string `json:"name"`
	//UserId     int64  `json:"user_id"`
	//CreateTime int64  `json:"create_time"`
}
