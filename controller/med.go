package controller

import (
	populate "PharmaProject/migration"
	model "PharmaProject/models"

	// "bufio"
	"errors"
	"fmt"

	// "os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

var Medlist = populate.MedFeed()

func medDb() []model.Medicine {
	return Medlist
}

func Printlist(meds []model.Medicine) {
	fmt.Println("The available medicines are : ")
	fmt.Printf("%s \n", strings.Repeat("-", 42))
	fmt.Printf("| %10s | %10s | %12s  |\n", "Id", "Name", "Price")
	fmt.Printf("%s \n", strings.Repeat("-", 42))
	for _, med := range meds {
		fmt.Printf("| %10d | %10s | %10d tk |\n", med.Id, med.Name, med.Price)
	}
	fmt.Printf("%s \n", strings.Repeat("-", 42))
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
	//
	// db.Create(&Medlist)

	result := db.Create(&M)
	if result.Error != nil {
		// fmt.Println(M.Id)
		return nil, result.Error
	}
	// fmt.Println(M.Id)
	return &M, nil
}

func (medicine *Medicine) DeleteMedicine(Id int) (bool, error) {
	// var med []Medicine
	// for i, medval := range Medlist {
	// 	if medval.Id == Id {
	// 		Medlist = append(Medlist[:i], Medlist[i+1:]...)
	// 		// fmt.Println(Medlist)
	// 		medfile, err := os.Create("./db/medicines.txt")
	// 		Check(err)

	// 		defer medfile.Close()
	// 		w := bufio.NewWriter(medfile)
	// 		for _, medval := range Medlist {
	// 			s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", medval.Id, medval.Name, medval.Price)
	// 			_, err := w.WriteString(s)
	// 			Check(err)
	// 		}
	// 		w.Flush()

	// 		return true, nil
	// 	}
	// }
	var med model.Medicine
	fmt.Println(Id)
	result := db.Delete(&med, Id) ///
	if result.RowsAffected > 0 {
		// fmt.Println(result.Error)
		return true, nil
	}
	return false, errors.New("Medicine does not exist")

}

func (medicine *Medicine) UpdateMedicine(med model.Medicine) (*model.Medicine, error) {
	// for i, medval := range Medlist {
	// 	if medval.Id == med.Id {
	// 		Medlist[i] = med
	// 		// fmt.Println(Medlist)
	// 		medfile, err := os.Create("./db/medicines.txt")
	// 		Check(err)

	// 		defer medfile.Close()
	// 		w := bufio.NewWriter(medfile)
	// 		for _, medval := range Medlist {
	// 			s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", medval.Id, medval.Name, medval.Price)
	// 			_, err := w.WriteString(s)
	// 			Check(err)
	// 		}
	// 		w.Flush()
	// 		return &med, nil
	// 	}
	// }
	var up model.Medicine
	result := db.First(&up, med.Id)
	if result.Error == nil {
		result = db.Model(&up).Updates(&med)
		// fmt.Println(result.Error)
		if result.RowsAffected > 0 {
			return &up, nil
		}
		return nil, result.Error
	}
	return nil, errors.New("Medicine does not exist")

}
