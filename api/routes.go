package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Init() {
	r := chi.NewRouter()
	r.Post("/login",Login)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", GetAllUsers)
		r.Get("/{userid}", GetUserByID)
		r.Post("/add", AddUser)
		r.Delete("/delete/{userid}", DeleteUserbyID)
		r.Post("/update/{userid}", UpdateUserbyID)
	})

	r.Route("/medicine", func(r chi.Router) {
		r.Get("/", GetAllMedicines)
		r.Get("/{med_id}", GetMedbyID)
		r.Post("/add", AddMedicine)
		r.Delete("/delete/{med_id}", DeleteMedicinebyID)
		r.Post("/update/{med_id}", UpdateMedicinebyID)

	})
	r.Route("/cart", func(r chi.Router) {
		r.Get("/", GetAllfromCart)
		r.Get("/{med_id}", GetItemfromCart)
		r.Post("/add", AddtoCart)
		r.Delete("/delete/{med_id}", RemovefromCartbyID)

	})

	r.Route("/order", func(r chi.Router) {
		r.Get("/", GetAllOrder)
		r.Post("/confirm", ConfirmOrder)

	})
	port := ":3000"
	fmt.Println("Listening to port ", port)
	http.ListenAndServe(port, r)

}
