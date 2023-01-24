package controller

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type User struct {
	Username string
	Password string
	Role string
}

var Userlist = Users()

func Users() []User {
	user := make([]User, 0)

	userfile, err := os.Open("./db/users.txt")
	check(err)

	defer userfile.Close()

	var lines []string
	scanner := bufio.NewScanner(userfile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, val := range lines {
		arr := strings.Split(val, ", ")
		var Username string
		var Password string
		var Role string
		for _, val2 := range arr {
			u := strings.Split(val2, ": ")
			if u[0] == "Username" {
				Username= u[1]
			}
			if u[0] == "Password" {
				Password = u[1]
			}
			if u[0] == "Role" {
				Role = u[1]
			}
		}
		user = append(user, User{
			Username: Username,
			Password: Password,
			Role : Role,
		})
	}
	return user
}

func Login(username, password string) (*User, error) {
	for i:=range Userlist{
		if Userlist[i].Username==username && Userlist[i].Password==password {
			return &Userlist[i],nil
		}
	}
	return nil,errors.New("Username and Password does not match")
}
