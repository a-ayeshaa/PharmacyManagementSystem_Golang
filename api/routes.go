package api

import (
	con "PharmaProject/api-controller"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func Init() {
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Get("/", con.GetAllUsers)
		r.Get("/{userid}", con.GetUserByID)
		r.Post("/add", con.AddUser)
		r.Get("/delete/{userid}", con.DeleteUserbyID)
		r.Post("/update/{userid}", con.UpdateUserbyID)
	})

	r.Route("/medicine",func(r chi.Router) {
		r.Get("/",con.GetAllMedicines)
	})
	port := ":3000"
	fmt.Println("Listening to port ", port)
	http.ListenAndServe(port, r)

}
