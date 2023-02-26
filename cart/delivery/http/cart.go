package http

import (
	"PharmaProject/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi/v5"
)

type CartHandler struct {
	CartUseCase domain.CartUseCase
}

// New will initialize the resources endpoint
func New(r *chi.Mux, cartUseCase domain.CartUseCase) {
	handler := &CartHandler{
		CartUseCase: cartUseCase,
	}

	r.Route("/cart", func(r chi.Router) {
		r.Get("/", handler.GetAllfromCart)
		r.Get("/{med_id}", handler.GetItemfromCart)
		r.Post("/add", handler.AddtoCart)
		r.Delete("/delete/{med_id}", handler.RemovefromCartbyID)

	})
}

func (h *CartHandler) GetAllfromCart(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allCart := h.CartUseCase.GetAllfromCart()
	err := json.NewEncoder(response).Encode(allCart)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CartHandler) GetItemfromCart(response http.ResponseWriter, request *http.Request) {
	medid, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	item, err := h.CartUseCase.GetItemfromCart(medid)
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

func (h *CartHandler) AddtoCart(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var cart domain.Cart

	err := json.NewDecoder(request.Body).Decode(&cart)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := govalidator.ValidateStruct(cart)
	if err != nil {
		// println("error: " + err.Error())
		http.Error(response, err.Error(), http.StatusBadRequest)
		return

	}
	println(result)
	newmed, err := h.CartUseCase.AddtoCart(cart)
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

func (h *CartHandler) RemovefromCartbyID(response http.ResponseWriter, request *http.Request) {
	med_id, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	med, err := h.CartUseCase.RemovefromCart(med_id)
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
