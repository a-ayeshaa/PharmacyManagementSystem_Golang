package api

import (
	con "PharmaProject/controller"
	model "PharmaProject/models"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func GetAllOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allMeds := con.NewOrder().GetAllOrder()
	err := json.NewEncoder(response).Encode(allMeds)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConfirmOrder(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var order model.Order
	err := json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		// fmt.Println("//12//")
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := govalidator.ValidateStruct(order)
	if err != nil {
		// println("error: " + err.Error())
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}
	println(result)
	neworder, err := con.NewOrder().ConfirmOrder(order.Username)
	if err != nil {
		// fmt.Println("//123//")
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&order)

	err = json.NewEncoder(response).Encode(neworder)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
