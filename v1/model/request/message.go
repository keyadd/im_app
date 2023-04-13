package request

type Message struct {
	FriendId int64  `json:"friend_id"`
	Content  string `json:"content"`
}

type RoomMessage struct {
	GroupNo int64  `json:"group_no"`
	Content string `json:"content"`
}

type AddMessage struct {
	Id          int64 `json:"id"`
	RecipientId int64 `json:"recipient_id"`
	SenderId    int64 `json:"sender_id"`
	Content     int64 `json:"content"`
	FileId      int64 `json:"file_id"`
	Type        int64 `json:"type"`
	Status      int64 `json:"status"`
	CreateTime  int64 `json:"create_time"`
	MessageType int64 `json:"message_type"`
}
