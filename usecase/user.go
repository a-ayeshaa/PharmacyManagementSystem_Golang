package controller

import (
	database "PharmaProject/db"
	model "PharmaProject/models"
	repo "PharmaProject/repository"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Role     string
}

var db = database.Connect()

func NewUser() UserController {
	return &User{}
}
func (u *User) Register(user model.RegisterUser) (*model.User, error) {
	if user.Password == user.Confirm_password {
		u := model.User{
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
			Password: user.Password,
		}
		return repo.NewUserRepo().Register(u)
	}

	return nil, errors.New("Confirm Password does not match with Password")
}

func (u *User) GetAllUsers() []model.User {
	return repo.NewUserRepo().GetAllUsers()
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

func (u *User) Login(login model.Login) (*model.User, error) {
	return repo.NewUserRepo().Login(login)
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	return repo.NewUserRepo().GetUserByID(id)
}

func (u *User) DeleteUserbyID(id int) (bool, error) {
	return repo.NewUserRepo().DeleteUserbyID(id)

}

func (u *User) UpdateUserbyID(user model.User) (*model.User, error) {
	result, err := repo.NewUserRepo().GetUserByID(user.ID)
	if err == nil {
		return repo.NewUserRepo().UpdateUser(*result, user)
	}

	return nil, err
}
