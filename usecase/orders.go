package controller

import (
	model "PharmaProject/models"
	"PharmaProject/repository"
	"errors"
)

type Order struct {
	Id         int
	Username   string
	Totalprice int
}

func NewOrder() OrderController {
	return &Order{}
}

func (or *Order) GetAllOrder() []model.Order {
	return repository.NewOrderRepo().GetAllOrder()
}

func AddOrder(o *model.Order) *model.Order {
	return repository.NewOrderRepo().AddOrder(o)

}

func (or *Order) ConfirmOrder(username string) (*model.Order, error) {
	// var cart []model.Cart
	// db.Find(&cart)
	cart:=repository.NewCartRepo().GetAllfromCart()
	if len(cart) != 0 {
		var total int = 0
		for _, val := range cart {
			total += val.Totalprice
		}

		order := model.Order{
			Username:   username,
			Totalprice: total,
		}
		newo := AddOrder(&order)
		return newo, nil
	}

	return nil, errors.New("Cart is empty!")
}
