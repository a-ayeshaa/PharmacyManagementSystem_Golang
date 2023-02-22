package api

import (
	worker "PharmaProject/worker"
	model "PharmaProject/models"
	con "PharmaProject/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/go-chi/chi/v5"
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

func GetMedbyID(response http.ResponseWriter, request *http.Request) {
	medid, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	user, err := con.NewMedicine().GetMedicine(medid)
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

func AddMedicine(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var med model.Medicine
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
	newmed, err := con.NewMedicine().AddMedicine(med)
	if err != nil {
		// println("error: " + err.Error())
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

func AddBulkMedicine(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	jsonVal:=request.Body
	worker.SendTask(jsonVal)
	go func() {
		worker.ReceiveTask()
	}()
}

func DeleteMedicinebyID(response http.ResponseWriter, request *http.Request) {
	med_id, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	med, err := con.NewMedicine().DeleteMedicine(med_id)
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

func UpdateMedicinebyID(response http.ResponseWriter, request *http.Request) {
	med_id, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var update_med model.Medicine
	update_med.Id = med_id
	err := json.NewDecoder(request.Body).Decode(&update_med)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	update_med.Id = med_id
	// fmt.Println(updateuser.ID)
	user, err := con.NewMedicine().UpdateMedicine(update_med)
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
