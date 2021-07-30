package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"simple-payment-api/app/database"
)

type App struct {
	Router *mux.Router
	DB     *database.DB
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	app.initRoutes()
	fmt.Println("Payment service listening...")
	return app
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/payment/", app.CreatePaymentHandler()).Methods("POST")
	app.Router.HandleFunc("/payment/{id}", app.PaymentStatusHandler()).Methods("GET")
}
