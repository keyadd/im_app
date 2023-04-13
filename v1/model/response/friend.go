package response

type FriendList struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	UserFriendId int64  `json:"user_friend_id"`
	CreateTime   int64  `json:"create_time"`
	User         []User `json:"user" gorm:"foreignKey:id;"`
}
