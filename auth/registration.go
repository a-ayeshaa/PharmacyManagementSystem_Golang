package auth

import (
	con "PharmaProject/controller"
	"fmt"
)

func HandleValidation() {
	if r := recover(); r != nil {
		fmt.Println("\n Re-enter details\n", r)
		Registration()
	}
}
func Registration() {
	var username, password, confpassword, role, email string
	fmt.Print("Enter Username: ")
	fmt.Scanln(&username)
	defer HandleValidation()
	err := con.ValidateUser(username)
	con.Check(err)
	fmt.Print("Enter Email: ")
	fmt.Scanln(&email)
	err = con.ValidateUser(email)
	con.Check(err)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)
	fmt.Print("Re-type Password: ")
	fmt.Scanln(&confpassword)
	fmt.Print("Enter Role: ")
	fmt.Scanln(&role)
}
