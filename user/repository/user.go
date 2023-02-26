package repository

import (
	"PharmaProject/domain"
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct{
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &User{
		db: db,
	}
}
func (u *User) Register(user domain.User) (*domain.User, error) {

	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil

}

func (u *User) GetAllUsers() []domain.User {
	var users []domain.User
	u.db.Find(&users)
	return users
}

func (u *User) Login(login domain.Login) (*domain.User, error) {
	var user domain.User
	result := u.db.First(&user, &login)
	if result.Error == nil {
		return &user, nil
	}
	return nil, errors.New("Username and Password does not match\n")
}

func (u *User) GetUserByID(id int) (*domain.User, error) {
	var user domain.User
	result := u.db.First(&user, id)
	if result.Error == nil {
		return &user, nil
	}
	return nil, errors.New("User does not exist!")
}

func (u *User) DeleteUserbyID(id int) (bool, error) {
	var user domain.User
	result := u.db.Delete(&user, id)
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, errors.New("User does not exist")

}

func (u *User) UpdateUser(user domain.User, updated_user domain.User) (*domain.User, error) {
	result := u.db.Model(&user).Updates(&updated_user)
	if result.RowsAffected > 0 {
		return &updated_user, nil
	}
	return nil, result.Error

}
