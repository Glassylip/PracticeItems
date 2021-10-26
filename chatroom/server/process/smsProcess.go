package process2

import (
	"fmt"
	"moudle/common/message"
	"moudle/server/utils"
	"net"

	"encoding/json"
)

type SmsProcess struct {
}

// Forward message
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	// Traverse the onlineUsers map[int]*UserProcess on the server side
	// Forward the message and take it out
	// Get the content of the SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		// If you need to filter to yourself, don't send it to yourself
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn, //
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("Failed to forward message err=", err)
	}
}
