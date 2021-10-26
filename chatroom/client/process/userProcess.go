package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"moudle/client/utils"
	"moudle/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int,
	userPwd string, userName string) (err error) {

	//1. Connect to server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	defer conn.Close()

	//2. Ready to send messages to the service via conn
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3. Create a LoginMes structure
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.Serialize registerMes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5. Assign data to the mes.Data field
	mes.Data = string(data)

	// 6. Serialize mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// Transfer instance
	tf := &utils.Transfer{
		Conn: conn,
	}

	// Send data to server
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Register to send information error err=", err)
	}

	mes, err = tf.ReadPkg() // mes is RegisterResMes

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	// Deserialize the Data part of mes into RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("Registered successfully, please log in again")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//1. connect to server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//2. Ready to send messages to the service via conn
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. Create a LoginMes structure
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. Serialize loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5. Assign data to the mes.Data field
	mes.Data = string(data)

	// 6. Serialize mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 7. Data is the message we want to send
	// 7.1 First send the length of data to the server
	// First get the length of data -> convert it into a byte cut that represents the length
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// Send length
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	fmt.Printf("Client, the length of the message sent=%d 内容=%s", len(data), string(data))

	// Send message
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	// Process the message returned by the server.
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	// Deserialize the Data part of the mes into LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("The list of current online users is as follows:")
		for _, v := range loginResMes.UsersId {
			// filter myself
			if v == userId {
				continue
			}

			fmt.Println("User id:\t", v)
			// Complete the client's onlineUsers complete initialization
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		// Start a coroutine on the client
		// The coroutine maintains communication with the server. If the server has data pushed to the client
		// Then it is received and displayed on the client terminal.
		go serverProcessMes(conn)

		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
