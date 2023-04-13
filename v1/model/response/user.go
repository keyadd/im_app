package response

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey;column:id;" `
	Password   string `json:"password"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Age        int64  `json:"age"`
	DeleteTime int64  `json:"delete_time"`
	Nickname   string `json:"nickname"`
	Online     int64  `json:"online"`
	CreateTime int64  `json:"create_time"`
	AvatarId   int64  `json:"avatar_id"`
	UpdateTime int64  `json:"update_time"`
}

type UserList struct {
	Id       int64  `json:"id" gorm:"primaryKey;column:id;" `
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int64  `json:"age"`
	//DeleteTime int64  `json:"delete_time"`
	Nickname string `json:"nickname"`
	Online   int64  `json:"online"`
	//CreateTime int64  `json:"create_time"`
	AvatarId int64 `json:"avatar_id"`
	//UpdateTime int64  `json:"update_time"`
}
