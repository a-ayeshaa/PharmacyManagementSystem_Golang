package repository

import (
	"PharmaProject/domain"

	"github.com/jinzhu/gorm"
)

type Order struct {
	db       *gorm.DB
	cartRepo domain.CartRepository
}

func New(db *gorm.DB) domain.OrderRepository {
	return &Order{
		db: db,
	}
}

func (or *Order) GetAllOrder() []domain.Order {
	var order []domain.Order
	or.db.Find(&order)
	return order
}

func (or *Order) AddOrder(o *domain.Order) *domain.Order {
	or.db.Create(&o)
	or.cartRepo.EmptyCart()
	return o

}
