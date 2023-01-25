package controller

import(
	model "PharmaProject/models"

)

type Controller interface {
	GetAll() []model.Medicine
	Get(Id int) model.Medicine
	Add(med model.Medicine) model.Medicine
	Delete(Id int) bool
	Update(med model.Medicine) model.Medicine
}
