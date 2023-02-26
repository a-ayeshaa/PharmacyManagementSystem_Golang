package usecase

import (
	"PharmaProject/domain"
	"errors"
	"net/mail"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type User struct {
	userRepo domain.UserRepository
}

func New(userRepo domain.UserRepository) domain.UserUseCase {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) Register(user domain.RegisterUser) (*domain.User, error) {
	if user.Password == user.Confirm_password {
		ureg := domain.User{
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
			Password: user.Password,
		}
		return u.userRepo.Register(ureg)
	}

	return nil, errors.New("Confirm Password does not match with Password")
}

func (u *User) GetAllUsers() []domain.User {
	return u.userRepo.GetAllUsers()
}

func Validate(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("Email must contain @address.com \n e.g your-name@gmail.com")

	}
	return nil
}

func (u *User) Login(login domain.Login) (*domain.User, error) {
	return u.userRepo.Login(login)
}

func (u *User) GetUserByID(id int) (*domain.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *User) DeleteUserbyID(id int) (bool, error) {
	return u.userRepo.DeleteUserbyID(id)

}

func (u *User) UpdateUserbyID(user domain.User) (*domain.User, error) {
	result, err := u.GetUserByID(user.ID)
	if err == nil {
		return u.userRepo.UpdateUser(*result, user)
	}

	return nil, err
}
