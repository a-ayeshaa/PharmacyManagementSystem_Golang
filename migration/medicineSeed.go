package migration

// import (
// 	model "PharmaProject/models"
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func Check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// //Seeding
// func medSeed() []model.Medicine {
// 	med := make([]model.Medicine, 10)
// 	for i := 0; i < 10; i++ {
// 		var m model.Medicine
// 		m.Id = i
// 		m.Name = "Napa" + strconv.Itoa(i)
// 		m.Price = rand.Intn(1000-500) + 500
// 		med[i] = m
// 	}

// 	medfile, err := os.OpenFile("./db/medicines.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	Check(err)

// 	defer medfile.Close()
// 	w := bufio.NewWriter(medfile)
// 	for _, medval := range med {
// 		s := fmt.Sprintf("ID: %d, Name: %s, Price: %d \n", medval.Id, medval.Name, medval.Price)
// 		_, err := w.WriteString(s)
// 		Check(err)
// 	}
// 	w.Flush()
// 	return med
// }

// func MedFeed() []model.Medicine {
// 	med := make([]model.Medicine, 0)

// 	medfile, err := os.Open("./db/medicines.txt")
// 	Check(err)

// 	defer medfile.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(medfile)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	for _, val := range lines {
// 		arr := strings.Split(val, ", ")
// 		var Id int
// 		var Name string
// 		var Price int
// 		for _, val2 := range arr {
// 			valmed := strings.Split(val2, ": ")
// 			if valmed[0] == "ID" {
// 				Id, _ = strconv.Atoi(valmed[1])
// 			}
// 			if valmed[0] == "Name" {
// 				Name = valmed[1]
// 			}
// 			if valmed[0] == "Price" {
// 				Price, _ = strconv.Atoi(strings.TrimSpace(valmed[1]))
// 				// fmt.Println(Price,valmed[1])
// 			}
// 		}
// 		med = append(med, model.Medicine{
// 			Id:    Id,
// 			Name:  Name,
// 			Price: Price,
// 		})
// 	}
// 	// fmt.Println(med)
// 	// fmt.Println(len(med))

// 	return med
// }