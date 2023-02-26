package http

import (
	"PharmaProject/domain"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct{
	userUseCase	domain.UserUseCase
}

func New(r *chi.Mux,userUseCase domain.UserUseCase){
	handler:=&UserHandler{
		userUseCase: userUseCase,
	}

	r.Post("/login", handler.Login)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
		r.Get("/{userid}", handler.GetUserByID)
		r.Post("/add", handler.AddUser)
		r.Delete("/delete/{userid}", handler.DeleteUserbyID)
		r.Post("/update/{userid}", handler.UpdateUserbyID)
	})
}

func(h *UserHandler) Login(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var user domain.Login
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
	newUser, err :=h.userUseCase.Login(user)
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

func(h *UserHandler) GetAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allUsers := h.userUseCase.GetAllUsers()
	err := json.NewEncoder(response).Encode(allUsers)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func(h *UserHandler) GetUserByID(response http.ResponseWriter, request *http.Request) {
	userid, _ := strconv.Atoi(chi.URLParam(request, "userid"))
	user, err := h.userUseCase.GetUserByID(userid)
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

func(h *UserHandler) AddUser(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var user domain.RegisterUser
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
	newUser, err := h.userUseCase.Register(user)
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

func(h *UserHandler) DeleteUserbyID(response http.ResponseWriter, request *http.Request) {
	userid, _ := strconv.Atoi(chi.URLParam(request, "userid"))
	user, err := h.userUseCase.DeleteUserbyID(userid)
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

func(h *UserHandler) UpdateUserbyID(response http.ResponseWriter, request *http.Request) {
	userid, _ := strconv.Atoi(chi.URLParam(request, "userid"))
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var updateuser domain.User
	err := json.NewDecoder(request.Body).Decode(&updateuser)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	updateuser.ID = userid
	// fmt.Println(updateuser.ID)
	user, err := h.userUseCase.UpdateUserbyID(updateuser)
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
