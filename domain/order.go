package domain

// Order ...
type Order struct {
	Id         int    `json:"id"`
	Username   string `json:"username" valid:"required"`
	Totalprice int    `json:"total_price" `
}

// OrderUseCase ...
type OrderUseCase interface {
	ConfirmOrder(username string) (*Order,error)
	GetAllOrder() []Order
}

// OrderRepository ...
type OrderRepository interface {
	AddOrder(*Order) *Order
	GetAllOrder() []Order
}

