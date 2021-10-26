package main

import (
	"fmt"
	"moudle/client/process"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {

	// get user input
	var key int

	for {
		fmt.Println("----------------Welcome to ChatRoom------------")
		fmt.Println("\t\t\t 1 Login ChatRoom")
		fmt.Println("\t\t\t 2 Register")
		fmt.Println("\t\t\t 3 Exit")
		fmt.Println("\t\t\t Please select(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Login")
			fmt.Println("Input User ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Input Password")
			fmt.Scanf("%s\n", &userPwd)
			// 1. login
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("Register")
			fmt.Println("Input User ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("Input Password")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("Input nickname:")
			fmt.Scanf("%s\n", &userName)
			//2. regist
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("Exit")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("Your input is wrong, please re-enter")
		}

	}

}
