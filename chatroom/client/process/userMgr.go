package process

import (
	"fmt"
	"moudle/client/model"
	"moudle/common/message"
)

// The map to be maintained by the client
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser // After the user logs in successfully, complete the initialization of CurUser

// Show online users in client
func outputOnlineUser() {
	fmt.Println("List of current online users:")
	for id, _ := range onlineUsers {
		fmt.Println("User id:\t", id)
	}
}

// Process the returned NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
