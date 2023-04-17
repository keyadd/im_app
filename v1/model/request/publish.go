package request

type Jsep struct {
	Sdp  any    `json:"sdp"`
	Type string `json:"type"`
}

type Publish struct {
	Jsep     Jsep  `json:"jsep"`
	FriendId int64 `json:"friend_id"`
}
