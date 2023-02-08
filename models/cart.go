package models

type Cart struct {
	Id         int        `json:"id" gorm:"primaryKey"`
	MedicineId int        `json:"medicine_id" valid:"required"`
	Name       string     `json:"name"`
	Totalprice int        `json:"total_price"`
	Quantity   int        `json:"quantity" valid:"required"`
	// Medicines  []Medicine `gorm:"foreignKey:Id;references:MedicineId"`
}
