package models

type Order struct {
	Id         int    `json:"id"`
	Username   string `json:"username" valid:"required"`
	Totalprice int    `json:"total_price" `
}
