package request

type GroupRequest struct {
	Name string `json:"name"`
}

type CreateGroup struct {
	GroupNo    int64  `json:"group_no"`
	Name       string `json:"name"`
	UserId     int64  `json:"user_id"`
	CreateTime int64  `json:"create_time"`
}

type Group struct {
	GroupNo    int64  `json:"group_no"`
	Name       string `json:"name"`
	UserId     int64  `json:"user_id"`
	CreateTime int64  `json:"create_time"`
}

type JoinGroup struct {
	GroupNo int64 `json:"group_no"`
}
