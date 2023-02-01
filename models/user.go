package models

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"-" binding:"required"`
	Role            string `json:"role"`
}
