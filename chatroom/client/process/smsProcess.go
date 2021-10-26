package process

import (
	"encoding/json"
	"fmt"
	"moudle/client/utils"
	"moudle/common/message"
)

type SmsProcess struct {
}

// Send group message
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	// SmsMes Instance
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// Serialize smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	// Serialize mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	// Send mes to server
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}

	return
}
