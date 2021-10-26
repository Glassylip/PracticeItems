package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"moudle/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {

	//1. Connect to server
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

	// 7. data is the message we want to send
	// 7.1 First send the length of data to the server
	// Get the length of data first -> convert it into a byte slice representing the length
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// Send length
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	// Send the message itself
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	// Process the message returned by the server.
	mes, err = readPkg(conn)

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	// Deserialize the Data part of the mes into LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("login successful")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
