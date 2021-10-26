package process2

import (
	"encoding/json"
	"fmt"
	"moudle/common/message"
	"moudle/server/model"
	"moudle/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// Notify other online users, I'm online
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	// Traverse onlineUsers, and then send NotifyUserStatusMes one by one
	for id, up := range userMgr.onlineUsers {
		// Filter to yourself
		if id == userId {
			continue
		}
		// Start notification [Write a method separately]
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// Serialize notifyUserStatusMes
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// Assign the serialized notifyUserStatusMes to mes.Data
	mes.Data = string(data)

	// Serialize the mes again and prepare to send
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// Create our Transfer instance
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//1.First take out mes.Data from the mes and directly deserialize it into RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// Go to the redis database to complete the verification.
	//1.Use model.MyUserDao to redis to verify
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "Unknown error in registration..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4. Assign data to resMes
	resMes.Data = string(data)

	//5. Serialize resMes and prepare to send
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6. Send data, we encapsulate it in the writePkg function
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

// Processing login request
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1. First take out mes.Data from the mes and directly deserialize it into LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	// Go to the redis database to complete the verification.
	//1.Use model.MyUserDao to redis to verify
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "Server internal error..."
		}

	} else {
		loginResMes.Code = 200
		// Because the user login is successful, we put the successful login into userMgr
		// Assign the userId of the user who successfully logged in to this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		// Notify other online users that I am online
		this.NotifyOthersOnlineUser(loginMes.UserId)
		// Put the id of the current online user into loginResMes.UsersId
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "login successful")
	}

	//3 Serialize loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4. Assign data to resMes
	resMes.Data = string(data)

	//5. Serialize resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6. Send data, we encapsulate it in the write Pkg function
	// Because of the use of hierarchical mode (mvc), we first create a Transfer instance, and then read
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
