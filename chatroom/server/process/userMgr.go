package process2

import (
	"fmt"
)

//Because there is only one UserMgr instance on the server side
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// Complete the initialization of userMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// Finish adding onlineUsers
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// Delete
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// Return all current online users
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// Return the corresponding value according to id
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	up, ok := this.onlineUsers[userId]
	if !ok { // The user you are looking for is currently offline.
		err = fmt.Errorf("User%d dose not exit", userId)
		return
	}
	return
}
