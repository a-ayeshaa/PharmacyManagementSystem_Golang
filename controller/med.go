package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	model "PharmaProject/models"
	populate "PharmaProject/migration"
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

func NewMedicine() Controller {
	return &Medicine{}
}
// var medicine model.Medicine

func (medicine *Medicine) GetAll() []model.Medicine {
	// fmt.Println("Hello world", 1234)
	return medDb()
}

func (medicine *Medicine) Get(Id int) model.Medicine {
	meds := medDb()

	return meds[Id]
}

func (medicine *Medicine) Add(M model.Medicine) model.Medicine {
	Medlist = append(Medlist, M)
	medfile, err := os.OpenFile("./db/medicines.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	//adding to file
	defer medfile.Close()
	w := bufio.NewWriter(medfile)
	s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", M.Id, M.Name, M.Price)
	_, err1 := w.WriteString(s)
	Check(err1)
	w.Flush()
	return M
}

func (medicine *Medicine) Delete(Id int) bool {
	// var med []Medicine
	for i, medval := range Medlist {
		if medval.Id == Id {
			Medlist = append(Medlist[:i], Medlist[i+1:]...)
			fmt.Println(Medlist)
		}
	}

	medfile, err := os.Create("./db/medicines.txt")
	Check(err)

	defer medfile.Close()
	w := bufio.NewWriter(medfile)
	for _, medval := range Medlist {
		s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", medval.Id, medval.Name, medval.Price)
		_, err := w.WriteString(s)
		Check(err)
	}
	w.Flush()

	return true
}

func (medicine *Medicine) Update(med model.Medicine) model.Medicine {
	for i, medval := range Medlist {
		if medval.Id == med.Id {
			Medlist[i] = med
		}
	}

	fmt.Println(Medlist)
	medfile, err := os.Create("./db/medicines.txt")
	Check(err)

	defer medfile.Close()
	w := bufio.NewWriter(medfile)
	for _, medval := range Medlist {
		s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", medval.Id, medval.Name, medval.Price)
		_, err := w.WriteString(s)
		Check(err)
	}
	w.Flush()
	return med
}
