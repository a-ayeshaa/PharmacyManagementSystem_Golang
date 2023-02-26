package repository

import (
	"PharmaProject/domain"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Medicine struct{
	db *gorm.DB
}

func New(db *gorm.DB) domain.MedicineRepository {
	return &Medicine{
		db: db,
	}
}

func (medicine *Medicine) GetAllMedicines() []domain.Medicine {
	var med []domain.Medicine
	medicine.db.Find(&med)
	return med
}

func (medicine *Medicine) GetMedicine(Id int) (*domain.Medicine, error) {
	var med domain.Medicine
	result := medicine.db.First(&med, Id)
	if result.Error == nil {
		return &med, nil
	}
	return nil, errors.New("Medicine could not be found")
}

func (medicine *Medicine) AddMedicine(M domain.Medicine) (*domain.Medicine, error) {
	fmt.Println(M)
	result := medicine.db.Create(&M)
	if result.Error != nil {
		return nil, result.Error
	}
	return &M, nil
}

func (medicine *Medicine) AddBulkMedicine(M []domain.Medicine) (*[]domain.Medicine, error) {
	fmt.Println(M)
	for _,val:= range M{
		result := medicine.db.Create(&val)
		if result.Error != nil {
			return nil, result.Error
		}
		
	}
	
	return &M, nil
}
func (medicine *Medicine) DeleteMedicine(Id int) (bool, error) {
	var med domain.Medicine
	fmt.Println(Id)
	result := medicine.db.Delete(&med, Id) ///
	if result.RowsAffected > 0 {
		// fmt.Println(result.Error)
		return true, nil
	}
	return false, errors.New("Medicine does not exist")

}

func (medicine *Medicine) UpdateMedicine(med domain.Medicine, update_med domain.Medicine) (*domain.Medicine, error) {

	result := medicine.db.Model(&med).Updates(&update_med)
	if result.RowsAffected > 0 {
		return &med, nil
	}
	return nil, result.Error

}
