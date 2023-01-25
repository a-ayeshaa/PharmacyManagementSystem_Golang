package controller

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type User struct {
	Username string
	Email    string
	Password string
	Role     string
}

var Userlist = Users()

func Register(username, password, confpassword, email, role string) (*User, error) {
	if password == confpassword {
		user := User{
			Username: username,
			Password: password,
			Role:     role,
			Email:    email,
		}
		Userlist = append(Userlist, user)
		return &user, nil
	}

	return nil, errors.New("Confirm Password does not match with Password")
}

func Users() []User {
	user := make([]User, 0)

	userfile, err := os.Open("./db/users.txt")
	Check(err)

	defer userfile.Close()

	var lines []string
	scanner := bufio.NewScanner(userfile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, val := range lines {
		arr := strings.Split(val, ", ")
		var new User
		for _, val2 := range arr {
			u := strings.Split(val2, ": ")
			// fmt.Println("u:",u)
			if u[0] == "Username" {
				new.Username = u[1]
			}
			if u[0] == "Password" {
				new.Password = u[1]
			}
			if u[0] == "Role" {
				new.Role = u[1]
			}
			if u[0] == "Email" {
				new.Email = u[1]
			}
		}
		user = append(user, new)
	}
	return user
}

func ValidateUser(val string) error {
	for _, user := range Userlist {
		if user.Username == val || user.Email == val {
			return errors.New("User already exists\n")
		}
	}
	return nil
}

func Login(username, password string) (*User, error) {
	for i := range Userlist {
		if Userlist[i].Username == username && Userlist[i].Password == password {
			return &Userlist[i], nil
		}
	}
	return nil, errors.New("Username and Password does not match\n")
}
