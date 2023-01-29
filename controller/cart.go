package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	model "PharmaProject/models"
)

type Cart struct {
	Id         int
	Name       string
	Totalprice int
	Quantity   int
}

var Cartlist []model.Cart

func Cartlists() []model.Cart {
	return Cartlist
}

func NewCart() CartController{
	return &Cart{}
}

func (ca *Cart) GetAllfromCart() []model.Cart{
	return Cartlist
}


func (ca *Cart) AddtoCart(c model.Cart) model.Cart {
	for i := range Cartlist {
		if Cartlist[i].Id == c.Id {
			Cartlist[i].Quantity += c.Quantity
			Cartlist[i].Totalprice += c.Totalprice
			fmt.Println(Cartlist)
			cartfile, err := os.Create("./db/carts.txt")
			Check(err)

			defer cartfile.Close()
			w := bufio.NewWriter(cartfile)
			for _, ca := range Cartlist {
				s := fmt.Sprintf("ID: %d, Name: %s, TotalPrice: %d, Quantity: %d \n", ca.Id, ca.Name, ca.Totalprice, ca.Quantity)
				_, err := w.WriteString(s)
				Check(err)
			}
			w.Flush()
			return c
		}
	}
	Cartlist = append(Cartlist, c)
	cartfile, err := os.OpenFile("./db/carts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	defer cartfile.Close()
	w := bufio.NewWriter(cartfile)
	s := fmt.Sprintf("ID: %d, Name: %s, TotalPrice: %d, Quantity: %d \n", c.Id, c.Name, c.Totalprice, c.Quantity)
	_, err1 := w.WriteString(s)
	Check(err1)
	w.Flush()
	return c
}

func (ca *Cart) RemovefromCart(id int) bool {
	for i, cartval := range Cartlist {
		if cartval.Id == id {
			Cartlist = append(Cartlist[:i], Cartlist[i+1:]...)
			break
		}

	}
	// fmt.Println(Cartlist)
	cartfile, err := os.Create("./db/carts.txt")
	Check(err)

	defer cartfile.Close()
	w := bufio.NewWriter(cartfile)
	for _, c := range Cartlist {
		s := fmt.Sprintf("ID: %d, Name: %s, TotalPrice: %d, Quantity: %d \n", c.Id, c.Name, c.Totalprice, c.Quantity)
		_, err := w.WriteString(s)
		Check(err)
	}
	w.Flush()
	return true
}

func (ca *Cart) PrintCart(carts []model.Cart) {
	if len(carts) == 0 {
		fmt.Println("Cart is empty")
	} else {
		var total int = 0
		fmt.Println("Your shopping cart : ")
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		fmt.Printf("| %10s | %10s | %12s  |\n", "Name", "Quantity", "Price")
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		for _, Cartlist := range Cartlist {
			total += Cartlist.Totalprice
			fmt.Printf("| %10s | %10d | %10d tk |\n", Cartlist.Name, Cartlist.Quantity, Cartlist.Totalprice)
		}
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		fmt.Printf("Total Price : %d tk \n", total)
		fmt.Printf("%s \n", strings.Repeat("-", 42))
	}
}

func (ca *Cart) ConfirmOrder(username string) bool {
	if len(Cartlist) != 0 {
		var total int = 0
		for _, val := range Cartlist {
			total += val.Totalprice
		}

		o := newOrder()
		order := model.Order{
			Username:   username,
			Totalprice: total,
		}
		fmt.Println(o.AddOrder(order))
		return true
	}
	return false
}

func SearchMed(id int) model.Cart {
	m := Medicine{}
	val := m.GetMedicine(id)
	newcart := model.Cart{
		Id:         val.Id,
		Name:       val.Name,
		Totalprice: val.Price,
		Quantity:   0,
	}
	return newcart
}
