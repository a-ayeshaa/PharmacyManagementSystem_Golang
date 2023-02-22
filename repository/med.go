package repository

import (
	model "PharmaProject/models"
	"errors"
	"fmt"
)

type Medicine struct {
	Id    int
	Name  string
	Price int
}
type MedicineRepo interface {
	GetAllMedicines() []model.Medicine
	GetMedicine(Id int) (*model.Medicine, error)
	AddMedicine(med model.Medicine) (*model.Medicine, error)
	AddBulkMedicine(meds []model.Medicine) (*[]model.Medicine,error)
	DeleteMedicine(Id int) (bool, error)
	UpdateMedicine(med model.Medicine, update_med model.Medicine) (*model.Medicine, error)
}

func NewMedicineRepo() MedicineRepo {
	return &Medicine{}
}

func (medicine *Medicine) GetAllMedicines() []model.Medicine {
	var med []model.Medicine
	db.Find(&med)
	return med
}

func (medicine *Medicine) GetMedicine(Id int) (*model.Medicine, error) {
	var med model.Medicine
	result := db.First(&med, Id)
	if result.Error == nil {
		return &med, nil
	}
	return nil, errors.New("Medicine could not be found")
}

func (medicine *Medicine) AddMedicine(M model.Medicine) (*model.Medicine, error) {
	fmt.Println(M)
	result := db.Create(&M)
	if result.Error != nil {
		return nil, result.Error
	}
	return &M, nil
}

func (medicine *Medicine) AddBulkMedicine(M []model.Medicine) (*[]model.Medicine, error) {
	// fmt.Println(M)
	result := db.Create(&M)
	if result.Error != nil {
		return nil, result.Error
	}
	return &M, nil
}
func (medicine *Medicine) DeleteMedicine(Id int) (bool, error) {
	var med model.Medicine
	fmt.Println(Id)
	result := db.Delete(&med, Id) ///
	if result.RowsAffected > 0 {
		// fmt.Println(result.Error)
		return true, nil
	}
	return false, errors.New("Medicine does not exist")

}

func (medicine *Medicine) UpdateMedicine(med model.Medicine, update_med model.Medicine) (*model.Medicine, error) {

	result := db.Model(&med).Updates(&update_med)
	if result.RowsAffected > 0 {
		return &update_med, nil
	}
	return nil, result.Error

}
