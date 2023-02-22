package controller

import (
	model "PharmaProject/models"
	"PharmaProject/repository"
	"errors"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Medicine struct {
	Id    int
	Name  string
	Price int
}

func NewMedicine() MedicineController {
	return &Medicine{}
}

func (medicine *Medicine) GetAllMedicines() []model.Medicine {
	return repository.NewMedicineRepo().GetAllMedicines()
}

func (medicine *Medicine) GetMedicine(Id int) (*model.Medicine, error) {
	return repository.NewMedicineRepo().GetMedicine(Id)
}

func (medicine *Medicine) AddMedicine(M model.Medicine) (*model.Medicine, error) {
	return repository.NewMedicineRepo().AddMedicine(M)
}
func (medicine *Medicine) AddBulkMedicine(M []model.Medicine) (*[]model.Medicine, error) {
	return repository.NewMedicineRepo().AddBulkMedicine(M)
}

func (medicine *Medicine) DeleteMedicine(Id int) (bool, error) {
	return repository.NewMedicineRepo().DeleteMedicine(Id)

}

func (medicine *Medicine) UpdateMedicine(med model.Medicine) (*model.Medicine, error) {
	result, err := repository.NewMedicineRepo().GetMedicine(med.Id)
	if err == nil {
		return repository.NewMedicineRepo().UpdateMedicine(*result, med)
	}
	return nil, errors.New("Medicine does not exist")

}
