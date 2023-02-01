package auth

import (
	con "PharmaProject/controller"
	model "PharmaProject/models"
	"fmt"
	"os"
	"strings"
)

func HandleValidationLogin() {
	if r := recover(); r != nil {
		fmt.Println("\n Re-enter details\n", r)
		Login()
	}
}
func Login() {
	defer HandleValidationLogin()

	m := con.NewMedicine()
	c := con.Cart{}
	meds := m.GetAllMedicines()
	newUser:=con.NewUser()

	var username string
	var password string
	fmt.Print("Enter Username:	")
	fmt.Scanln(&username)
	fmt.Print("Enter Password:	")
	fmt.Scanln(&password)

	user, err := newUser.Login(username, password)
	// fmt.Println(user)
	if user != nil && err == nil {
		con.Printlist(meds)
		if strings.ToLower(user.Role) == "admin" {
			var end bool = false
			for !end {
				var choice int
				fmt.Println("Admin, you have the following options to execute: \n 1. Add Medicine \n 2. Delete Medicine \n 3. Edit Medicine \n 4. Get Medicine \n 5. Exit")
				fmt.Println("Choose an option: ")
				fmt.Scanln(&choice)
				switch choice {
				case 1:
					var price int
					var name string
					fmt.Println("Set Name: ")
					fmt.Scanln(&name)
					fmt.Println("Set Price: ")
					fmt.Scanln(&price)

					newmed := model.Medicine{
						// Id:    index,
						Name:  name,
						Price: price,
					}
					fmt.Println(m.AddMedicine(newmed))
					con.Printlist(m.GetAllMedicines())

				case 2:
					var index int
					fmt.Println("Enter the ID of the medicine you want to delete: ")
					fmt.Scanln(&index)
					m.DeleteMedicine(index)
					con.Printlist(m.GetAllMedicines())

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

					medupdate := model.Medicine{
						Id:    index,
						Name:  name,
						Price: price,
					}
					fmt.Println(m.UpdateMedicine(medupdate))
					con.Printlist(m.GetAllMedicines())

				case 4:
					var index int
					fmt.Println("Enter the Id of the medicine :")
					fmt.Scanln(&index)
					fmt.Println(m.GetMedicine(index))
				case 5:
					// err := os.Remove("./db/carts.txt")
					// if err!=nil{
					// 	panic(err)
					// }
					end = true
				}

			}
		} else if strings.ToLower(user.Role) == "customer" {
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
					c.RemovefromCart(id)
					fmt.Println(con.Cartlists())
				case 3:
					c.PrintCart(con.Cartlist)
				case 4:
					val := c.ConfirmOrder(user.Username)
					if val {
						fmt.Println("Your order has been confirmed, order again!")
					} else {
						fmt.Println("Your cart is empty.")
					}
				case 5:
					err := os.Truncate("./db/carts.txt", 0)
					if err != nil {
						panic(err)
					}
					end = true
				}
			}
		}
	} else {
		con.Check(err)
	}

}
