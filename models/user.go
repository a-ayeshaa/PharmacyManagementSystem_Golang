package models

import "time"

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

type User struct {
	ID        int    `json:"id" gorm:"primaryKey" `
	Username  string `json:"username" valid:"required" gorm:"not null;unique"`
	Email     string `json:"email" valid:"required" gorm:"not null;unique"`
	Password  string `json:"password" valid:"required" gorm:"not null"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
