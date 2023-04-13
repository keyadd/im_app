package request

type AddFriend struct {
	FriendId int64 `json:"friend_id"`
}

type Friend struct {
	UserId       int64 `json:"user_id"`
	UserFriendId int64 `json:"user_friend_id"`
	CreateTime   int64 `json:"create_time"`
}
