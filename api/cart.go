package api

import (
	con "PharmaProject/controller"
	model "PharmaProject/models"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/asaskevich/govalidator"

	"github.com/go-chi/chi/v5"
)

func GetAllfromCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allCart := con.NewCart().GetAllfromCart()
	err := json.NewEncoder(response).Encode(allCart)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetItemfromCart(response http.ResponseWriter, request *http.Request) {
	medid, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	item, err := con.NewCart().GetItemfromCart(medid)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(item)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddtoCart(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var med model.Cart
	
	err := json.NewDecoder(request.Body).Decode(&med)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := govalidator.ValidateStruct(med)
	if err != nil {
		// println("error: " + err.Error())
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}
	println(result)
	newmed,err := con.NewCart().AddtoCart(med)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(response).Encode(newmed)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RemovefromCartbyID(response http.ResponseWriter, request *http.Request) {
	med_id, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	med, err := con.NewCart().RemovefromCart(med_id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(med)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}