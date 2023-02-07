package controller

import (
	model "PharmaProject/models"
)

type UserController interface {
	GetAllUsers() []model.User
	RegisterUser(username, password, confpassword, email, role string) (*model.User, error)
	Register(model.RegisterUser) (*model.User, error)
	Login(username, password string) (*model.User, error)
	ValidateUser(val string) error
	GetUserByID(id int) (*model.User, error)
	DeleteUserbyID(id int) (bool, error)
	UpdateUserbyID(model.User) (*model.User, error)
}

type MedicineController interface {
	GetAllMedicines() []model.Medicine
	GetMedicine(Id int) (*model.Medicine, error)
	AddMedicine(med model.Medicine) model.Medicine
	DeleteMedicine(Id int) (bool, error)
	UpdateMedicine(med model.Medicine) (*model.Medicine, error)
}

type CartController interface {
	GetAllfromCart() []model.Cart
	AddtoCart(cart model.Cart) (*model.Cart,error)
	GetItemfromCart(int) (*model.Cart,error)
	RemovefromCart(Id int) (bool,error)
}

type OrderController interface {
	ConfirmOrder(username string) (*model.Order,error)
	GetAllOrder() []model.Order
}

// type SnackController interface{

// }
