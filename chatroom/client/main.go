package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {

	// Get user input
	var key int
	var loop = true

	for loop {
		fmt.Println("----------------Welcome to log in to the multi-person chat system------------")
		fmt.Println("\t\t\t 1 Log in to the chat room")
		fmt.Println("\t\t\t 2 Registered user")
		fmt.Println("\t\t\t 3 Exit system")
		fmt.Println("\t\t\t Please choose(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Log in to the chat room")
			loop = false
		case 2:
			fmt.Println("Registered user")
			loop = false
		case 3:
			fmt.Println("Exit system")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("Your input is wrong, please re-enter")
		}

	}
	// login
	if key == 1 {
		fmt.Println("Please enter the id of the user")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("Please enter the user's password")
		fmt.Scanf("%s\n", &userPwd)

		login(userId, userPwd)

	} else if key == 2 {
		fmt.Println("Logic for user registration....")
	}
}
