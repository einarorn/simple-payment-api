package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	app.Router.Handle("/login/", app.LoginHandler()).Methods(http.MethodPost)
	app.Router.Handle("/payment/", authMiddleware(app.CreatePaymentHandler())).Methods(http.MethodPost)
	app.Router.Handle("/payment/{id}", authMiddleware(app.PaymentStatusHandler())).Methods(http.MethodGet)
}
