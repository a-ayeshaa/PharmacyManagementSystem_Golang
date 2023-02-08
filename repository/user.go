package repository

import (
	model "PharmaProject/models"
	"errors"
)

type UserRepo interface {
	GetAllUsers() []model.User
	Register(model.User) (*model.User, error)
	Login(model.Login) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	DeleteUserbyID(id int) (bool, error)
	UpdateUser(user model.User, updated_user model.User) (*model.User, error)
}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Role     string
}

func NewUserRepo() UserRepo {
	return &User{}
}

func (u *User) Register(user model.User) (*model.User, error) {

	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil

}

func (u *User) GetAllUsers() []model.User {
	var users []model.User
	db.Find(&users)
	return users
}

func (u *User) Login(login model.Login) (*model.User, error) {
	var user model.User
	result := db.First(&user, &login)
	if result.Error == nil {
		return &user, nil
	}
	return nil, errors.New("Username and Password does not match\n")
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	var user model.User
	result := db.First(&user, id)
	if result.Error == nil {
		return &user, nil
	}
	return nil, errors.New("User does not exist!")
}

func (u *User) DeleteUserbyID(id int) (bool, error) {
	var user model.User
	result := db.Delete(&user, id)
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, errors.New("User does not exist")

}

func (u *User) UpdateUser(user model.User, updated_user model.User) (*model.User, error) {
	result := db.Model(&user).Updates(&updated_user)
	if result.RowsAffected > 0 {
		return &updated_user, nil
	}
	return nil, result.Error

}
