package http

import (
	"PharmaProject/domain"
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	OrderUseCase domain.OrderUseCase
}

// New will initialize the resources endpoint
func New(r *chi.Mux, orderUseCase domain.OrderUseCase) {
	handler := &OrderHandler{
		OrderUseCase: orderUseCase,
	}

	r.Route("/order", func(r chi.Router) {
		r.Get("/", handler.GetAllOrder)
		r.Post("/confirm", handler.ConfirmOrder)

	})
}

func (h *OrderHandler) GetAllOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allMeds := h.OrderUseCase.GetAllOrder()
	err := json.NewEncoder(response).Encode(allMeds)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) ConfirmOrder(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var order domain.Order
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
	neworder, err := h.OrderUseCase.ConfirmOrder(order.Username)
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
