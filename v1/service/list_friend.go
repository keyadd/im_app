package service

import (
	"fmt"
	"im_app/v1/model"
	"im_app/v1/model/response"
)

func GetFriendList(id any) (list []response.UserList) {
	user_id := id.(int64)
	u, err := model.GetFriendList(user_id)
	fmt.Println(err)

	var l []response.UserList
	for _, list := range u {
		for _, user := range list.User {
			userList := response.UserList{
				Id:       user.Id,
				Username: user.Username,
				Email:    user.Email,
				Nickname: user.Nickname,
				AvatarId: user.AvatarId,
			}
			l = append(l, userList)
		}
	}
	return l

}
