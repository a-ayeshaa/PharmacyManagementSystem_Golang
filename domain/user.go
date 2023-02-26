package domain

import "time"

// RegisterUser ...
type RegisterUser struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	Username         string `json:"username" valid:"required" gorm:"not null" gorm:"unique"`
	Email            string `json:"email" valid:"required" gorm:"not null" gorm:"unique"`
	Password         string `json:"password" valid:"required" gorm:"not null"`
	Confirm_password string `json:"confirm_password" valid:"required" gorm:"-:migration" gorm:"not null"`
	Role             string `json:"role"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// User ...
type User struct {
	ID        int    `json:"id" gorm:"primaryKey" `
	Username  string `json:"username" valid:"required" gorm:"not null;unique"`
	Email     string `json:"email" valid:"required" gorm:"not null;unique"`
	Password  string `json:"password" valid:"required" gorm:"not null"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


// Login ...
type Login struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// UserUseCase  ...
type UserUseCase interface {
	GetAllUsers() []User
	Register(RegisterUser) (*User, error)
	Login(Login) (*User, error)
	GetUserByID(id int) (*User, error)
	DeleteUserbyID(id int) (bool, error)
	UpdateUserbyID(User) (*User, error)
}


// UserRepository
type UserRepository interface {
	GetAllUsers() []User
	Register(User) (*User, error)
	Login(Login) (*User, error)
	GetUserByID(id int) (*User, error)
	DeleteUserbyID(id int) (bool, error)
	UpdateUser(user User, updated_user User) (*User, error)
}