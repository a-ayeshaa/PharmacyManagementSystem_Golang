package cmd

import (
	// api "PharmaProject/api"
	cartHttp "PharmaProject/cart/delivery/http"
	cartRepo "PharmaProject/cart/repository"
	cartUseCase "PharmaProject/cart/usecase"
	"PharmaProject/internal/conn"
	medicineHttp "PharmaProject/medicine/delivery/http"
	medRepo "PharmaProject/medicine/repository"
	medUseCase "PharmaProject/medicine/usecase"
	orderHttp "PharmaProject/order/delivery/http"
	orderRepo "PharmaProject/order/repository"
	orderUseCase "PharmaProject/order/usecase"
	userHttp "PharmaProject/user/delivery/http"
	userRepo "PharmaProject/user/repository"
	userUseCase "PharmaProject/user/usecase"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
)

var startAPICmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s", "ser", "serv"},
	// SuggestFor: []string{"start"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// config.init()

	},
	Short: "Initiates the Pharmacy Management System",
	Run: func(cmd *cobra.Command, args []string) {
		r := chi.NewRouter()

		db := conn.ConnectDB()
		//initialise the repositories
		userRepo := userRepo.New(db)
		medicineRepo := medRepo.New(db)
		cartRepo := cartRepo.New(db)
		orderRepo := orderRepo.New(db)

		//initialise the use cases
		userUseCase := userUseCase.New(userRepo)
		medicineUseCase := medUseCase.New(medicineRepo)
		cartUseCase := cartUseCase.New(cartRepo, medicineRepo)
		orderUseCase := orderUseCase.New(orderRepo, cartRepo)

		//initialise the handlers
		userHttp.New(r, userUseCase)
		orderHttp.New(r, orderUseCase)
		medicineHttp.New(r, medicineUseCase)
		cartHttp.New(r, cartUseCase)

		// r := api.Init()
		if err := conn.ConnectWorker(); err != nil {
			fmt.Printf("failed to connect worker: %v", err)
		}
		if err := conn.ConnectRedis(); err != nil {
			panic(fmt.Errorf("Can't connect redis: %v", err))
		}
		port := ":3000"
		fmt.Println("Listening to port ", port)
		http.ListenAndServe(port, r)
	},
}

func init() {
	rootCmd.AddCommand(startAPICmd)
}
