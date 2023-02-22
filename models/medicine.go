package models

type Medicine struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" valid:"required" `
	Price int    `json:"price" valid:"required"`
}
