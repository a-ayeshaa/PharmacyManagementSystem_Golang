package repository

import (
	model "PharmaProject/models"
	"errors"
	"fmt"
)

type Cart struct {
	Id         int
	MedicineId int
	Name       string
	Totalprice int
	Quantity   int
}
type CartRepo interface {
	GetAllfromCart() []model.Cart
	AddtoCart(cart model.Cart) (*model.Cart, error)
	GetItemfromCart(int) (*model.Cart, error)
	RemovefromCart(Id int) (bool, error)
	EmptyCart()
	GetfromCart(model.Cart) (*model.Cart,error)
	UpdateCart(model.Cart,model.Cart) (*model.Cart,error)
}

func NewCartRepo() CartRepo {
	return &Cart{}
}

func (ca *Cart) GetfromCart(m model.Cart) (*model.Cart,error){
	var upmed model.Cart
	result := db.First(&upmed, m)
	return &upmed,result.Error
}

func (ca *Cart) GetAllfromCart() []model.Cart {
	var cart []model.Cart
	db.Find(&cart)
	return cart
}

func (ca *Cart) AddtoCart(c model.Cart) (*model.Cart, error) {
	db.Create(&c)
	return &c, nil
}

func (ca *Cart) GetItemfromCart(id int) (*model.Cart, error) {
	var med model.Cart
	result := db.First(&med, id)
	if result.Error == nil {
		return &med, nil
	}
	return nil, errors.New("Item with that Id could not be found")
}

func (ca *Cart) RemovefromCart(id int) (bool, error) {
	var med model.Cart
	fmt.Println(id)
	result := db.Delete(&med, id)
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, errors.New("Medicine could not be found")

}

func (ca *Cart) EmptyCart(){
	db.Exec("Truncate table carts")
}

func (ca *Cart) UpdateCart(cart model.Cart,update_cart model.Cart) (*model.Cart,error){
	result := db.Model(&cart).Updates(update_cart)
	if result.RowsAffected <= 0 {
		return nil, result.Error
	}
	return &cart, nil
}