package models

type Medicine struct {
	Id    int
	Name  string
	Price int
}

func NewMedicine() Product {
	return &Medicine{}
}

func (m *Medicine) SetID(id int) {
	m.Id = id
}

func (m *Medicine) SetName(name string) {

}
func (m *Medicine) SetPrice(price int) {

}
