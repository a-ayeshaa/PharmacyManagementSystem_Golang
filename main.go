package main

import (
	reg "PharmaProject/auth"
	"fmt"
)

func main() {
	for {
		var ch int
		fmt.Println("Welcome! You have the following options:\n 1. Login\n 2. Create a new account")
		fmt.Print("Enter option number: ")
		fmt.Scanln(&ch)
		switch ch {
		//LOGIN .....
		case 1:
			reg.Login()
			fmt.Println("testing")
		//REGISTRATION....
		case 2:
			reg.Registration()
		}
	}

}
