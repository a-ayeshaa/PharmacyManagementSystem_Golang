package apicontroller

import (
	con "PharmaProject/controller"
	// model "PharmaProject/models"
	"encoding/json"
	"net/http"
	// "strconv"
	// "github.com/go-chi/chi/v5"
)

func GetAllMedicines(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allMeds := con.NewMedicine().GetAllMedicines()
	err := json.NewEncoder(response).Encode(allMeds)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
