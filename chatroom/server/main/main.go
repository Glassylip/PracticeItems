package main

import (
	"fmt"
	"moudle/server/model"
	"net"
	"time"
)

// Handle communication with the client
func process(conn net.Conn) {
	defer conn.Close()

	// Call the master control here to create one
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("Client and server communication protocol error", err)
		return
	}
}

func init() {
	// When the server starts, we initialize our redis connection pool
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func initUserDao() {

	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	fmt.Println("The server [new structure] listens on port 8889....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	// Wait for the client to connect to the server
	for {
		fmt.Println("Wait for the client to connect to the server.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		// Start a coroutine to maintain communication with the client
		go process(conn)
	}
}
