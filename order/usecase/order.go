package usecase

import (
	"PharmaProject/domain"
	"errors"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Order struct {
	orderRepo domain.OrderRepository
	cartRepo  domain.CartRepository
}

func New(orderRepo domain.OrderRepository, cartRepo domain.CartRepository) domain.OrderUseCase {
	return &Order{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (or *Order) GetAllOrder() []domain.Order {
	return or.orderRepo.GetAllOrder()
}

func (or *Order) AddOrder(o *domain.Order) *domain.Order {
	return or.orderRepo.AddOrder(o)

}

func (or *Order) ConfirmOrder(username string) (*domain.Order, error) {
	// var cart []domain.Cart
	// db.Find(&cart)
	cart := or.cartRepo.GetAllfromCart()
	if len(cart) != 0 {
		var total int = 0
		for _, val := range cart {
			total += val.Totalprice
		}

		order := domain.Order{
			Username:   username,
			Totalprice: total,
		}
		newo := or.orderRepo.AddOrder(&order)
		return newo, nil
	}

	return nil, errors.New("Cart is empty!")
}
