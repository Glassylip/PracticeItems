package main

import (
	"fmt"
	"io"
	"moudle/common/message"
	process2 "moudle/server/process"
	"moudle/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// According to the different types of messages sent by the client, decide which function to call to process
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	fmt.Println("mes=", mes)

	switch mes.Type {
	// Login
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	// Register
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// Create an SmsProcess instance to complete forwarding group chat messages.
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("The message type does not exist and cannot be processed...")
	}
	return
}

func (this *Processor) process2() (err error) {

	// Send information cyclically
	for {
		// Read the data packet, encapsulate it into a function readPkg(), return Message, Err
		// Create a Transfer instance to complete the package reading task
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("The client exits, the server also exits..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}

		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}
