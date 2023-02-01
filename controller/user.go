package controller

import (
	model "PharmaProject/models"
	"bufio"
	"errors"
	"fmt"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Role     string
}

var Userlist = Users()

func NewUser() UserController {
	return &User{}
}

func (u *User) RegisterUser(username, password, confpassword, email, role string) (*model.User, error) {
	if password == confpassword {
		var index int = Userlist[len(Userlist)-1].ID + 1
		user := model.User{
			ID:       index,
			Username: username,
			Password: password,
			Role:     role,
			Email:    email,
		}
		Userlist = append(Userlist, user)

		userfile, err := os.OpenFile("./db/users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		Check(err)

		defer userfile.Close()
		w := bufio.NewWriter(userfile)
		s := fmt.Sprintf("Username: %s, Password: %s, Role: %s, Email: %s \n", user.Username, user.Password, user.Role, user.Email)
		_, err1 := w.WriteString(s)
		Check(err1)
		w.Flush()
		return &user, nil
	}

	return nil, errors.New("Confirm Password does not match with Password")
}
func (u *User) Register(user model.User) (*model.User, error) {
	//Check if user already exists.....
	err := u.ValidateUser(user.Username)
	if err != nil {
		return nil, err
	}
	err = Validate(user.Email)
	if err != nil {
		return nil, err
	}
	u.ValidateUser(user.Email)
	if err != nil {
		return nil, err
	}
	if user.Password == user.ConfirmPassword {
		var index int = Userlist[len(Userlist)-1].ID + 1
		user.ID = index
		Userlist = append(Userlist, user)

		userfile, err := os.OpenFile("./db/users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		Check(err)

		defer userfile.Close()
		w := bufio.NewWriter(userfile)
		s := fmt.Sprintf("ID: %d,Username: %s, Password: %s, Role: %s, Email: %s \n", user.ID, user.Username, user.Password, user.Role, user.Email)
		_, err1 := w.WriteString(s)
		Check(err1)
		w.Flush()
		return &user, nil
	}

	return nil, errors.New("Confirm Password does not match with Password")
}

func Users() []model.User {
	user := make([]model.User, 0)

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
		var new model.User
		for _, val2 := range arr {
			u := strings.Split(val2, ": ")
			// fmt.Println("u:",u)
			if u[0] == "Username" {
				new.Username = u[1]
			}
			if u[0] == "ID" {
				new.ID, _ = strconv.Atoi(u[1])
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

func (u *User) GetAllUsers() []model.User {
	return Userlist
}

func (u *User) ValidateUser(val string) error {
	for _, user := range Userlist {
		if user.Username == val || user.Email == val {
			return errors.New("User already exists\n")
		}
	}
	return nil
}

func Validate(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("Email must contain @address.com \n e.g your-name@gmail.com")

	}
	return nil
}

func ValidatePass(pass string) {
	// rmatch, err := regexp.MatchString(`/^.*(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$ %^&*~><.,:;]).*$/i`,pass)
	rmatch, err := regexp.MatchString(`(?=abc)`, pass)
	fmt.Println(rmatch, err)

}

func (u *User) Login(username, password string) (*model.User, error) {
	for i := range Userlist {
		if Userlist[i].Username == username && Userlist[i].Password == password {
			return &Userlist[i], nil
		}
	}
	return nil, errors.New("Username and Password does not match\n")
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	for _, val := range Userlist {
		if val.ID == id {
			return &val, nil
		}

	}
	return nil, errors.New("User does not exist!")
}

func (u *User) DeleteUserbyID(id int) (bool, error) {
	for i, val := range Userlist {
		if val.ID == id {
			Userlist = append(Userlist[:i], Userlist[i+1:]...)
			// fmt.Println(Medlist)
			userfile, err := os.Create("./db/users.txt")
			Check(err)

			defer userfile.Close()
			w := bufio.NewWriter(userfile)
			for _, user := range Userlist {
				s := fmt.Sprintf("ID: %d,Username: %s, Password: %s, Role: %s, Email: %s \n", user.ID, user.Username, user.Password, user.Role, user.Email)
				_, err := w.WriteString(s)
				Check(err)
			}
			w.Flush()

			return true, nil
		}
	}

	return false, errors.New("User does not exist")

}

func (u *User)	UpdateUserbyID(user model.User) (*model.User,error){
	for i, val := range Userlist {
		if val.ID == user.ID {
			Userlist[i]=user
			userfile, err := os.Create("./db/users.txt")
			Check(err)
			fmt.Println(Userlist)
			defer userfile.Close()
			w := bufio.NewWriter(userfile)
			for _, user := range Userlist {
				s := fmt.Sprintf("ID: %d,Username: %s, Password: %s, Role: %s, Email: %s \n", user.ID, user.Username, user.Password, user.Role, user.Email)
				_, err := w.WriteString(s)
				Check(err)
			}
			w.Flush()

			return &Userlist[i], nil
		}
	}

	return nil,errors.New("User does not exist")
}