package models

type Cart struct {
	Id         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	Totalprice int    `json:"total_price"`
	Quantity   int    `json:"quantity" binding:"required"`
}
