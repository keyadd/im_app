package request

type UserGroupRequest struct {
	UserId     int64 `json:"user_id"`
	GroupId    int64 `json:"group_id"`
	CreateTime int64 `json:"create_time"`
}
