package controller

import (
	model "PharmaProject/models"
)

type UserController interface{
	GetAllUsers() []model.User
	RegisterUser(username, password, confpassword, email, role string) (*model.User, error)
	Register(model.User) (*model.User, error)
	Login(username, password string) (*model.User, error) 
	ValidateUser(val string) error
	GetUserByID(id int) (*model.User,error)
	DeleteUserbyID(id int) (bool,error)
	UpdateUserbyID(model.User) (*model.User,error)
}

type MedicineController interface {
	GetAllMedicines() []model.Medicine
	GetMedicine(Id int) model.Medicine
	AddMedicine(med model.Medicine) model.Medicine
	DeleteMedicine(Id int) bool
	UpdateMedicine(med model.Medicine) model.Medicine
}

type CartController interface{
	GetAllfromCart() []model.Cart
	AddtoCart(cart model.Cart) model.Cart
	PrintCart([]model.Cart)
	RemovefromCart(Id int) bool
	ConfirmOrder(username string) bool	
}

type OrderController interface{
	AddOrder(model.Order) model.Order
}

// type SnackController interface{

// }
