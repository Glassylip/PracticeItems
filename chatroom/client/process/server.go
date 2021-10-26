package process

import (
	"encoding/json"
	"fmt"
	"moudle/client/utils"
	"moudle/common/message"
	"net"
	"os"
)

// Show Menu after login
func ShowMenu() {

	fmt.Println("-------Login successful---------")
	fmt.Println("-------1. Show Online Users---------")
	fmt.Println("-------2. Send Message---------")
	fmt.Println("-------3. Message List---------")
	fmt.Println("-------4. Exit System---------")
	fmt.Println("Please select (1-4):")
	var key int
	var content string

	// SmsProcess instance
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("Input message:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("Message List")
	case 4:
		fmt.Println("Exit System...")
		os.Exit(0)
	default:
		fmt.Println("Wrong input..")
	}

}

// Keep communicating with the server
func serverProcessMes(conn net.Conn) {
	// Create a transfer instance and read the messages sent by the server continuously
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("The client is waiting to read the message sent by the server")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		// Message read
		switch mes.Type {

		case message.NotifyUserStatusMesType: // Someone is online

			//1. Get.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2. Save the user's information and state to the customer map[int]User
			updateUserStatus(&notifyUserStatusMes)

		case message.SmsMesType: // Message from the crowd
			outputGroupMes(&mes)
		default:
			fmt.Println("The server returned an unknown message type")
		}

	}
}
