package repository

import (
	"PharmaProject/domain"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Cart struct{
	db *gorm.DB
}

func New(db *gorm.DB) domain.CartRepository {
	return &Cart{
		db: db,
	}
}

func (ca *Cart) GetfromCart(m domain.Cart) (*domain.Cart, error) {
	var upmed domain.Cart
	result := ca.db.First(&upmed, m)
	return &upmed, result.Error
}

func (ca *Cart) GetAllfromCart() []domain.Cart {
	var cart []domain.Cart
	ca.db.Find(&cart)
	return cart
}

func (ca *Cart) AddtoCart(c domain.Cart) (*domain.Cart, error) {
	ca.db.Create(&c)
	return &c, nil
}

func (ca *Cart) GetItemfromCart(id int) (*domain.Cart, error) {
	var med domain.Cart
	result := ca.db.First(&med, id)
	if result.Error == nil {
		return &med, nil
	}
	return nil, errors.New("Item with that Id could not be found")
}

func (ca *Cart) RemovefromCart(id int) (bool, error) {
	var med domain.Cart
	fmt.Println(id)
	result := ca.db.Delete(&med, id)
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, errors.New("Medicine could not be found")

}

func (ca *Cart) EmptyCart() {
	ca.db.Exec("Truncate table carts")
}

func (ca *Cart) UpdateCart(cart domain.Cart, update_cart domain.Cart) (*domain.Cart, error) {
	result := ca.db.Model(&cart).Updates(update_cart)
	if result.RowsAffected <= 0 {
		return nil, result.Error
	}
	return &cart, nil
}
