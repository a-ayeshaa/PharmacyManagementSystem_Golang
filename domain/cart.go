package domain

// Cart ...
type Cart struct {
	Id         int        `json:"id" gorm:"primaryKey"`
	MedicineId int        `json:"medicine_id" valid:"required"`
	Name       string     `json:"name"`
	Totalprice int        `json:"total_price"`
	Quantity   int        `json:"quantity" valid:"required"`
	// Medicines  []Medicine `gorm:"foreignKey:Id;references:MedicineId"`
}

// CartRepository ...
type CartRepository interface {
	GetAllfromCart() []Cart
	AddtoCart(cart Cart) (*Cart, error)
	GetItemfromCart(int) (*Cart, error)
	RemovefromCart(Id int) (bool, error)
	EmptyCart()
	GetfromCart(Cart) (*Cart, error)
	UpdateCart(Cart, Cart) (*Cart, error)
}

// CartUseCase ...
type CartUseCase interface {
	GetAllfromCart() []Cart
	AddtoCart(cart Cart) (*Cart,error)
	GetItemfromCart(int) (*Cart,error)
	RemovefromCart(Id int) (bool,error)
}