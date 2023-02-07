package models

type Cart struct {
	Id         int    `json:"id"`
	MedicineId int    `json:"med_id" valid:"required"`
	Name       string `json:"name"`
	Totalprice int    `json:"total_price"`
	Quantity   int    `json:"quantity" valid:"required"`
}
