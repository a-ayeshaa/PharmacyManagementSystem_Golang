package http

import (
	"PharmaProject/domain"
	"PharmaProject/internal/conn"
	"PharmaProject/task"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi/v5"
	"github.com/thedevsaddam/retry"
)

type MedicineHandler struct {
	MedicineUseCase	domain.MedicineUseCase
}

// New will initialize the resources endpoint
func New(r *chi.Mux,medUseCase domain.MedicineUseCase){
	handler:=&MedicineHandler{
		MedicineUseCase: medUseCase,
	}

	r.Route("/medicine", func(r chi.Router) {
		r.Get("/", handler.GetAllMedicines)
		r.Get("/{med_id}", handler.GetMedbyID)
		// r.Post("/add", AddMedicine)
		r.Post("/add", handler.AddBulkMedicine)
		r.Delete("/delete/{med_id}", handler.DeleteMedicinebyID)
		r.Post("/update/{med_id}", handler.UpdateMedicinebyID)

	})
}

func(h *MedicineHandler) GetAllMedicines(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	allMeds := h.MedicineUseCase.GetAllMedicines()
	err := json.NewEncoder(response).Encode(allMeds)
	response.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func(h *MedicineHandler) GetMedbyID(response http.ResponseWriter, request *http.Request) {
	medid, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	user, err := h.MedicineUseCase.GetMedicine(medid)
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

func(h *MedicineHandler) AddMedicine(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var med domain.Medicine
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
	newmed, err := h.MedicineUseCase.AddMedicine(med)
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

func(h *MedicineHandler) AddBulkMedicine(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	// jsonVal, err := ioutil.ReadAll(request.Body)
	var meds []domain.Medicine
	err := json.NewDecoder(request.Body).Decode(&meds)
	if err != nil {
		fmt.Println("Error while decoding")
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	for _, med := range meds {
		// fmt.Println(med)
		result, err := govalidator.ValidateStruct(med)
		if err != nil {
			// println("error: " + err.Error())
			http.Error(response, err.Error(), http.StatusBadRequest)
			return

		}
		println(result)
		if med.Price < 0 {
			http.Error(response, "Price must be greater than 0", http.StatusBadRequest)
			return
		}
	}
	pld,err := json.Marshal(meds)
	// json.NewEncoder(pld).Encode(meds)

	// fmt.Println("error")
	errCh := make(chan error, 1)
	go func() {
		err := retry.DoFunc(3, 1*time.Second, func() error {
			_, err := conn.DefaultWorker().SendTask(&tasks.Signature{
				Name:         task.TaskAddMedicine,
				RetryCount:   1,
				RetryTimeout: 10,
				Args: []tasks.Arg{
					{
						Type:  "string",
						Value: string(pld),
					},
				},
			})
			fmt.Println(err)
			return err
		})
		errCh <- err
	}()
	// worker.SendTask(jsonMeds.Bytes())
}

func(h *MedicineHandler) DeleteMedicinebyID(response http.ResponseWriter, request *http.Request) {
	med_id, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	med, err := h.MedicineUseCase.DeleteMedicine(med_id)
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

func(h *MedicineHandler) UpdateMedicinebyID(response http.ResponseWriter, request *http.Request) {
	med_id, _ := strconv.Atoi(chi.URLParam(request, "med_id"))
	contentType := request.Header.Get("Content-Type")
	if contentType != "" && contentType != "application/json" {
		http.Error(response, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	var update_med domain.Medicine
	update_med.Id = med_id
	err := json.NewDecoder(request.Body).Decode(&update_med)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	update_med.Id = med_id
	// fmt.Println(updateuser.ID)
	user, err := h.MedicineUseCase.UpdateMedicine(update_med)
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