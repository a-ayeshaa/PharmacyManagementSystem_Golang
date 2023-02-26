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

type Cart struct {
	cartRepo domain.CartRepository
	medicineRepo domain.MedicineRepository
}

func New(cartRepo domain.CartRepository, medicineRepo domain.MedicineRepository) domain.CartUseCase {
	return &Cart{
		cartRepo: cartRepo,
		medicineRepo: medicineRepo,
	}
}


func (ca *Cart) GetAllfromCart() []domain.Cart {
	return ca.cartRepo.GetAllfromCart()
}

func (ca *Cart) AddtoCart(c domain.Cart) (*domain.Cart, error) {
	newmedicine, err := ca.medicineRepo.GetMedicine(c.MedicineId)
	if err != nil {
		return nil, errors.New("Medicine does not exist")
	}
	c.Name = newmedicine.Name
	c.Totalprice = newmedicine.Price * c.Quantity

	upmed, err := ca.cartRepo.GetfromCart(domain.Cart{
		MedicineId: c.MedicineId,
	})
	if err != nil {
		return ca.cartRepo.AddtoCart(c)
	}
	return ca.cartRepo.UpdateCart(*upmed, domain.Cart{
		MedicineId: c.MedicineId,
		Totalprice: upmed.Totalprice + c.Totalprice,
		Quantity:   upmed.Quantity + c.Quantity,
	})
}

func (ca *Cart) GetItemfromCart(id int) (*domain.Cart, error) {
	return ca.cartRepo.GetItemfromCart(id)
}

func (ca *Cart) RemovefromCart(id int) (bool, error) {
	return ca.cartRepo.RemovefromCart(id)
}
