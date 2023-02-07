package controller

import (
	model "PharmaProject/models"
	// "bufio"
	"errors"
	"fmt"
	// "os"
	"strings"
)

type Cart struct {
	Id         int
	MedicineId int
	Name       string
	Totalprice int
	Quantity   int
}

var Cartlist []model.Cart

func Cartlists() []model.Cart {
	return Cartlist
}

func NewCart() CartController {
	return &Cart{}
}

func (ca *Cart) GetAllfromCart() []model.Cart {
	var cart []model.Cart
	db.Find(&cart)
	return cart
}

func (ca *Cart) AddtoCart(c model.Cart) (*model.Cart, error) {
	var newmedicine model.Medicine
	result := db.First(&newmedicine, c.MedicineId)
	if result.Error != nil {
		return nil, errors.New("Medicine does not exist")
	}
	c.Name = newmedicine.Name
	c.Totalprice = newmedicine.Price * c.Quantity
	var upmed model.Cart
	result = db.First(&upmed, &model.Cart{
		MedicineId: c.MedicineId,
	})
	if result.Error != nil {
		db.Create(&c)
		return &c, nil
	}
	result = db.Model(&upmed).Updates(&model.Cart{
		MedicineId: c.MedicineId,
		Totalprice: upmed.Totalprice + c.Totalprice,
		Quantity:   upmed.Quantity + c.Quantity,
	})
	if result.RowsAffected <= 0 {
		return nil, result.Error
	}
	return &upmed, nil
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
	result := db.Delete(&med, id) ///
	if result.RowsAffected > 0 {
		// fmt.Println(result.Error)
		return true, nil
	}
	return false, errors.New("Medicine could not be found")

}

func PrintCart(carts []model.Cart) {
	if len(carts) == 0 {
		fmt.Println("Cart is empty")
	} else {
		var total int = 0
		fmt.Println("Your shopping cart : ")
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		fmt.Printf("| %10s | %10s | %12s  |\n", "Name", "Quantity", "Price")
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		for _, Cartlist := range Cartlist {
			total += Cartlist.Totalprice
			fmt.Printf("| %10s | %10d | %10d tk |\n", Cartlist.Name, Cartlist.Quantity, Cartlist.Totalprice)
		}
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		fmt.Printf("Total Price : %d tk \n", total)
		fmt.Printf("%s \n", strings.Repeat("-", 42))
	}
}

func SearchMed(id int) (*model.Cart, error) {
	m := Medicine{}
	val, err := m.GetMedicine(id)
	if err != nil {
		return nil, err
	}
	newcart := model.Cart{
		Id:         val.Id,
		Name:       val.Name,
		Totalprice: val.Price,
		Quantity:   0,
	}
	return &newcart, nil
}
