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

func Login(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var user model.Login
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := govalidator.ValidateStruct(user)
	if err != nil {
		// println("error: " + err.Error())
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}
	println(result)
	newUser, err := con.NewUser().Login(user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusConflict)
	}
	err = json.NewEncoder(response).Encode(newUser)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allUsers := con.NewUser().GetAllUsers()
	err := json.NewEncoder(response).Encode(allUsers)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserByID(response http.ResponseWriter, request *http.Request) {
	userid, _ := strconv.Atoi(chi.URLParam(request, "userid"))
	user, err := con.NewUser().GetUserByID(userid)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(user)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddUser(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var user model.RegisterUser
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := govalidator.ValidateStruct(user)
	if err != nil {
		// println("error: " + err.Error())
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}
	println(result)
	newUser, err := con.NewUser().Register(user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusConflict)
	}
	err = json.NewEncoder(response).Encode(newUser)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteUserbyID(response http.ResponseWriter, request *http.Request) {
	userid, _ := strconv.Atoi(chi.URLParam(request, "userid"))
	user, err := con.NewUser().DeleteUserbyID(userid)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(user)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUserbyID(response http.ResponseWriter, request *http.Request) {
	userid, _ := strconv.Atoi(chi.URLParam(request, "userid"))
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var updateuser model.User
	err := json.NewDecoder(request.Body).Decode(&updateuser)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	updateuser.ID = userid
	// fmt.Println(updateuser.ID)
	user, err := con.NewUser().UpdateUserbyID(updateuser)
	if err != nil {
		http.Error(response, err.Error(), http.StatusConflict)
	}
	err = json.NewEncoder(response).Encode(user)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
