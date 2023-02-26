package domain

// Medicine ...
type Medicine struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" valid:"required" `
	Price int    `json:"price" valid:"required"`
}

// MedicineUseCase ....
type MedicineUseCase interface {
	GetAllMedicines() []Medicine
	GetMedicine(Id int) (*Medicine, error)
	AddMedicine(med Medicine) (*Medicine,error)
	AddBulkMedicine(meds []Medicine) (*[]Medicine,error)
	DeleteMedicine(Id int) (bool, error)
	UpdateMedicine(med Medicine) (*Medicine, error)
}


// MedicineRepository ....
type MedicineRepository interface{
	GetAllMedicines() []Medicine
	GetMedicine(Id int) (*Medicine, error)
	AddMedicine(med Medicine) (*Medicine, error)
	AddBulkMedicine(meds []Medicine) (*[]Medicine,error)
	DeleteMedicine(Id int) (bool, error)
	UpdateMedicine(med Medicine, update_med Medicine) (*Medicine, error)
}