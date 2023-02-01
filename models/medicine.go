package models

type Medicine struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

// func NewMedicine() Product {
// 	return &Medicine{}
// }

// func (m *Medicine) SetID(id int) {
// 	m.Id = id
// }

// func (m *Medicine) SetName(name string) {
// 	m.Name=name
// }
// func (m *Medicine) SetPrice(price int) {
// 	m.Price=price
// }
