package controller

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//Seeding
func medSeed() []Medicine {
	med := make([]Medicine, 10)
	for i := 0; i < 10; i++ {
		var m Medicine
		m.Id = i
		m.Name = "Napa" + strconv.Itoa(i)
		m.Price = rand.Intn(1000-500) + 500
		med[i] = m
	}
	
	medfile, err := os.OpenFile("./db/medicines.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	defer medfile.Close()
	w := bufio.NewWriter(medfile)
	for _, medval := range med {
		s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", medval.Id, medval.Name, medval.Price)
		_, err := w.WriteString(s)
		Check(err)
	}
	w.Flush()
	return med
}

func MedFeed() []Medicine{
	med := make([]Medicine, 0)

	medfile, err := os.Open("./db/medicines.txt")
    Check(err)

    defer medfile.Close()

    var lines []string
    scanner := bufio.NewScanner(medfile)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    for _,val:=range lines{
		arr:=strings.Split(val,", ")
		var Id int
		var Name string
		var Price int
		for _,val2:=range arr{
			valmed:=strings.Split(val2,": ")
			if valmed[0]=="ID"{
				Id,_=strconv.Atoi(valmed[1])
			}
			if valmed[0]=="Name"{
				Name=valmed[1]
			}
			if valmed[0]=="Price" {
				Price,_=strconv.Atoi(strings.TrimSpace(valmed[1]))
				// fmt.Println(Price,valmed[1])
			}
		}
		med = append(med, Medicine{
			Id: Id,
			Name: Name,
			Price: Price,
		})
	}
	// fmt.Println(med)
	// fmt.Println(len(med))

	return med
}

var Medlist = MedFeed()
// var MedLiist = MedFeed()

func medDb() []Medicine {
	return Medlist
}

func Printlist(meds []Medicine) {
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

type IMedicine interface {
	GetAll() []Medicine
	Get(Id int) Medicine
	Add(med Medicine) Medicine
	Delete(Id int) bool
	Update(med Medicine) Medicine
}

func (medicine Medicine) GetAll() []Medicine {
	return medDb()
}

func (medicine Medicine) Get(Id int) Medicine {
	meds := medDb()

	return meds[Id]
}

func (medicine Medicine) Add(M Medicine) Medicine {
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

func (medicine Medicine) Delete(Id int) bool {
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

func (medicine Medicine) Update(med Medicine) Medicine {
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
