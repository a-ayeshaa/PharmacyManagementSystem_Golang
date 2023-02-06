package models

type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username" valid:"required"`
	Email            string `json:"email" valid:"required"`
	Password         string `json:"password" valid:"required"`
	Confirm_password string `json:"-" valid:"required"`
	Role             string `json:"role"`
}
