package main

import (
	con "PharmaProject/controller"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(con.Medlist)
	m := con.Medicine{}
	c := con.Cart{}
	meds := m.GetAll()
	con.Printlist(meds)

	var user string
	fmt.Println("Enter Username:")
	fmt.Scanln(&user)
	if strings.ToLower(user) == "admin" {
		var end bool = false
		for !end {
			var choice int
			fmt.Println("Admin, you have the following options to execute: \n 1. Add Medicine \n 2. Delete Medicine \n 3. Edit Medicine \n 4. Get Medicine \n 5. Exit")
			fmt.Println("Choose an option: ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				var index int = con.Medlist[len(con.Medlist)-1].Id + 1
				var price int
				var name string
				fmt.Println("Set Name: ")
				fmt.Scanln(&name)
				fmt.Println("Set Price: ")
				fmt.Scanln(&price)

				newmed := con.Medicine{
					Id:    index,
					Name:  name,
					Price: price,
				}
				fmt.Println(m.Add(newmed))
				con.Printlist(m.GetAll())

			case 2:
				var index int
				fmt.Println("Enter the ID of the medicine you want to delete: ")
				fmt.Scanln(&index)
				m.Delete(index)
				con.Printlist(m.GetAll())

			case 3:
				var index int
				var price int
				var name string
				fmt.Println("Enter the Id of the medicine :")
				fmt.Scanln(&index)
				fmt.Println("Enter the Name of the medicine :")
				fmt.Scanln(&name)
				fmt.Println("Enter the Price of the medicine :")
				fmt.Scanln(&price)

				medupdate := con.Medicine{
					Id:    index,
					Name:  name,
					Price: price,
				}
				fmt.Println(m.Update(medupdate))
				con.Printlist(m.GetAll())

			case 4:
				var index int
				fmt.Println("Enter the Id of the medicine :")
				fmt.Scanln(&index)
				fmt.Println(m.Get(index))
			case 5:
				// err := os.Remove("./db/carts.txt")
				// if err!=nil{
				// 	panic(err)
				// }
				end = true
			}

		}
	} else if strings.ToLower(user) == "customer" {
		var end bool = false
		for !end {
			var choice int
			fmt.Println("Customer, you have the following options to execute: \n 1. Add Medicine to cart \n 2. Remove from Cart \n 3. Check Cart \n 4. Confirm Order \n 5. Exit")
			fmt.Println("Choose an option: ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				var id int
				var q int
				fmt.Println("Enter ID :")
				fmt.Scanln(&id)
				add := con.SearchMed(id)
				fmt.Println("Enter Quantity :")
				fmt.Scanln(&q)
				add.Totalprice = add.Totalprice * q
				add.Quantity = q
				// cart:=Cart{1,"Napa",200,2}
				c.AddtoCart(add)
				fmt.Println(con.Cartlists())
			case 2:
				var id int
				fmt.Println("Enter ID:")
				fmt.Scanln(&id)
				c.Remove(id)
				fmt.Println(con.Cartlists())
			case 3:
				c.Printcart(con.Cartlist)
			case 4:
				val := c.Confirm()
				if val {
					fmt.Println("Your order has been confirmed, order again!")
				} else {
					fmt.Println("Your cart is empty.")
				}
			case 5:
				err := os.Truncate("./db/carts.txt",0)
				if err != nil {
					panic(err)
				}
				end = true
			}
		}
	}

}
