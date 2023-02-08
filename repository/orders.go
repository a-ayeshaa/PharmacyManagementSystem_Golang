package repository

import (
	model "PharmaProject/models"
)

type Order struct {
	Id         int
	Username   string
	Totalprice int
}

type OrderRepo interface {
	AddOrder(*model.Order) *model.Order
	GetAllOrder() []model.Order
}

func NewOrderRepo() OrderRepo {
	return &Order{}
}

func (or *Order) GetAllOrder() []model.Order {
	var order []model.Order
	db.Find(&order)
	return order
}

func (or *Order) AddOrder(o *model.Order) *model.Order {
	db.Create(&o)
	NewCartRepo().EmptyCart()
	return o

}
