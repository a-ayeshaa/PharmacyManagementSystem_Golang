package controller

import (
	model "PharmaProject/models"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var OrderList = Orders()

type Order struct {
	Id         int
	Username   string
	Totalprice int
}

func Orders() []model.Order {
	order := make([]model.Order, 0)

	orderfile, err := os.Open("./db/orders.txt")
	Check(err)

	defer orderfile.Close()

	var lines []string
	scanner := bufio.NewScanner(orderfile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, val := range lines {
		arr := strings.Split(val, ", ")
		var new model.Order
		for _, val2 := range arr {
			u := strings.Split(val2, ": ")
			// fmt.Println("u:",u)
			if u[0] == "Username" {
				new.Username = u[1]
			}
			if u[0] == "ID" {
				new.Id, _ = strconv.Atoi(u[1])
				// fmt.Println(new.ID)
			}
			if u[0] == "Total Price" {
				new.Totalprice, _ = strconv.Atoi(u[1])
			}
		}
		order = append(order, new)
	}
	return order
}

func NewOrder() OrderController {
	return &Order{}
}

func (or *Order) GetAllOrder() []model.Order {
	var order []model.Order
	db.Find(&order)
	return order
}

func AddOrder(o *model.Order) *model.Order {
	// var id int
	// if len(OrderList) == 0 {
	// 	id = 0
	// } else {
	// 	id = OrderList[len(OrderList)-1].Id + 1
	// }
	// o.Id = id
	// OrderList = append(OrderList, o)

	// orderfile, err := os.OpenFile("./db/orders.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// Check(err)

	// //adding to file
	// defer orderfile.Close()
	// w := bufio.NewWriter(orderfile)
	// s := fmt.Sprintf("ID: %d, Username: %s, Total Price: %d \n", o.Id, o.Username, o.Totalprice)
	// _, err1 := w.WriteString(s)
	// Check(err1)
	// w.Flush()

	// Cartlist = make([]model.Cart, 0)
	// err2 := os.Truncate("./db/carts.txt", 0)
	// Check(err2)
	// return OrderList[id]
	db.Create(&o)
	db.Exec("Truncate table carts")
	return o

}

func (or *Order) ConfirmOrder(username string) (*model.Order, error) {
	var cart []model.Cart
	db.Find(&cart)
	fmt.Println(len(cart))
	if len(cart) != 0 {
		var total int = 0
		// fmt.Println("123")
		for _, val := range cart {
			total += val.Totalprice
		}

		order := model.Order{
			Username:   username,
			Totalprice: total,
		}
		// o,err := AddOrder(order)
		newo := AddOrder(&order)
		return newo, nil
	}
	// fmt.Println(len(cart)+1)

	return nil, errors.New("Cart is empty!")
}
