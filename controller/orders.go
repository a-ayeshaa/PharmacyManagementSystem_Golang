package controller

import (
	"bufio"
	"fmt"
	"os"
	model "PharmaProject/models"

)

var order []model.Order

type Order struct {
	Id         int
	Username   string
	Totalprice int
}

func newOrder() OrderController{
	return &Order{}
}

func (or *Order) AddOrder(o model.Order) model.Order {
	var id int
	if len(order) == 0 {
		id = 0
	} else {
		id = order[len(order)-1].Id + 1
	}
	o.Id = id
	order = append(order, o)

	orderfile, err := os.OpenFile("./db/orders.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	//adding to file
	defer orderfile.Close()
	w := bufio.NewWriter(orderfile)
	s := fmt.Sprintf("ID: %d, Username: %s, Total Price: %d \n", o.Id, o.Username, o.Totalprice)
	_, err1 := w.WriteString(s)
	Check(err1)
	w.Flush()

	Cartlist=make([]model.Cart, 0)
	err2:= os.Truncate("./db/carts.txt",0)
	Check(err2)
	return order[id]

}
