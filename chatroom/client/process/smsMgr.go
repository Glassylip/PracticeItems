package process

import (
	"encoding/json"
	"fmt"
	"moudle/common/message"
)

func outputGroupMes(mes *message.Message) {
	// Deserialize mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	// Show message
	info := fmt.Sprintf("user id:\t%d say:\t%s",
		smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
