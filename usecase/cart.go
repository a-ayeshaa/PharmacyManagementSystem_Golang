package controller

import (
	model "PharmaProject/models"
	"PharmaProject/repository"
	"errors"
)

type Cart struct {
	Id         int
	MedicineId int
	Name       string
	Totalprice int
	Quantity   int
}

func NewCart() CartController {
	return &Cart{}
}

func (ca *Cart) GetAllfromCart() []model.Cart {
	return repository.NewCartRepo().GetAllfromCart()
}

func (ca *Cart) AddtoCart(c model.Cart) (*model.Cart, error) {
	newmedicine, err := repository.NewMedicineRepo().GetMedicine(c.MedicineId)
	if err != nil {
		return nil, errors.New("Medicine does not exist")
	}
	c.Name = newmedicine.Name
	c.Totalprice = newmedicine.Price * c.Quantity

	upmed, err := repository.NewCartRepo().GetfromCart(model.Cart{
		MedicineId: c.MedicineId,
	})
	if err != nil {
		return repository.NewCartRepo().AddtoCart(c)
	}
	return repository.NewCartRepo().UpdateCart(*upmed, model.Cart{
		MedicineId: c.MedicineId,
		Totalprice: upmed.Totalprice + c.Totalprice,
		Quantity:   upmed.Quantity + c.Quantity,
	})
}

func (ca *Cart) GetItemfromCart(id int) (*model.Cart, error) {
	return repository.NewCartRepo().GetItemfromCart(id)
}

func (ca *Cart) RemovefromCart(id int) (bool, error) {
	return repository.NewCartRepo().RemovefromCart(id)
}
