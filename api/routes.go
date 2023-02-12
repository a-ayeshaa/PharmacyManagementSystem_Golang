package api

import (
	_ "fmt"
	_ "net/http"

	"github.com/go-chi/chi/v5"
)

func Init() *chi.Mux {
	r := chi.NewRouter()
	// r.Get("/",func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello ayesha!")
	// })
	r.Post("/login", Login)
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

	return r

}
