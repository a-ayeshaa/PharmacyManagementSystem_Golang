package usecase

import (
	"PharmaProject/domain"
	"errors"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Medicine struct {
	medicineRepo domain.MedicineRepository
}

func New(medicineRepo domain.MedicineRepository) domain.MedicineUseCase {
	return &Medicine{
		medicineRepo: medicineRepo,
	}
}

func (medicine *Medicine) GetAllMedicines() []domain.Medicine {
	return medicine.GetAllMedicines()
}

func (medicine *Medicine) GetMedicine(Id int) (*domain.Medicine, error) {
	return medicine.GetMedicine(Id)
}

func (medicine *Medicine) AddMedicine(med domain.Medicine) (*domain.Medicine, error) {
	return medicine.AddMedicine(med)
}
func (medicine *Medicine) AddBulkMedicine(meds []domain.Medicine) (*[]domain.Medicine, error) {
	return medicine.AddBulkMedicine(meds)
}

func (medicine *Medicine) DeleteMedicine(Id int) (bool, error) {
	return medicine.DeleteMedicine(Id)

}

func (medicine *Medicine) UpdateMedicine(med domain.Medicine) (*domain.Medicine, error) {
	result, err := medicine.GetMedicine(med.Id)
	if err == nil {
		return medicine.medicineRepo.UpdateMedicine(*result, med)
	}
	return nil, errors.New("Medicine does not exist")

}
