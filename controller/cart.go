package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cart struct {
	Id         int
	Name       string
	Totalprice int
	Quantity   int
}

var Cartlist []Cart

func Cartlists() []Cart {
	return Cartlist
}

type ICart interface {
	AddtoCart(c Cart) Cart
	Remove(id int) bool
	Printcart(carts []Cart)
	Confirm() bool
}

func (ca Cart) AddtoCart(c Cart) Cart {
	for i := range Cartlist {
		if Cartlist[i].Id == c.Id {
			Cartlist[i].Quantity += c.Quantity
			Cartlist[i].Totalprice += c.Totalprice
			fmt.Println(Cartlist)
			cartfile, err := os.Create("./db/carts.txt")
			check(err)

			defer cartfile.Close()
			w := bufio.NewWriter(cartfile)
			for _, ca := range Cartlist {
				s := fmt.Sprintf("ID: %d, Name: %s, TotalPrice: %d, Quantity: %d \n", ca.Id, ca.Name, ca.Totalprice, ca.Quantity)
				_, err := w.WriteString(s)
				check(err)
			}
			w.Flush()
			return c
		}
	}
	Cartlist = append(Cartlist, c)
	cartfile, err := os.OpenFile("./db/carts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer cartfile.Close()
	w := bufio.NewWriter(cartfile)
	s := fmt.Sprintf("ID: %d, Name: %s, TotalPrice: %d, Quantity: %d \n", c.Id, c.Name, c.Totalprice, c.Quantity)
	_, err1 := w.WriteString(s)
	check(err1)
	w.Flush()
	return c
}

func (ca Cart) Remove(id int) bool {
	for i, cartval := range Cartlist {
		if cartval.Id == id {
			Cartlist = append(Cartlist[:i], Cartlist[i+1:]...)
			break
		}

	}
	// fmt.Println(Cartlist)
	cartfile, err := os.Create("./db/carts.txt")
	check(err)

	defer cartfile.Close()
	w := bufio.NewWriter(cartfile)
	for _, c := range Cartlist {
		s := fmt.Sprintf("ID: %d, Name: %s, TotalPrice: %d, Quantity: %d \n", c.Id, c.Name, c.Totalprice, c.Quantity)
		_, err := w.WriteString(s)
		check(err)
	}
	w.Flush()
	return true
}

func (ca Cart) Printcart(carts []Cart) {
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

func (ca Cart) Confirm() bool {
	if len(Cartlist) != 0 {
		var total int = 0
		for _, val := range Cartlist {
			total += val.Totalprice
		}

		o := Order{}
		order := Order{
			Username:   "customer",
			Totalprice: total,
		}
		fmt.Println(o.AddOrder(order))
		return true
	}
	return false
}

func SearchMed(id int) Cart {
	m := Medicine{}
	val := m.Get(id)
	newcart := Cart{
		Id:         val.Id,
		Name:       val.Name,
		Totalprice: val.Price,
		Quantity:   0,
	}
	return newcart
}
