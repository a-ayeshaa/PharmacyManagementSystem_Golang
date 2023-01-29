package controller

import (
	model "PharmaProject/models"
)

// type ProductController interface {
// 	GetAllMedicines() []model.Medicine
// 	Get(Id int) model.Product
// 	Add(med model.Medicine) model.Medicine
// 	Delete(Id int) bool
// 	Update(med model.Medicine) model.Medicine
// }

type ProductController interface {
	GetAllMedicines() []model.Medicine
	GetMedicine(Id int) model.Medicine
	AddMedicine(med model.Medicine) model.Medicine
	DeleteMedicine(Id int) bool
	UpdateMedicine(med model.Medicine) model.Medicine
}
