package main

import (
	"log"
	"net/http"
	"simple-payment-api/app"
	"simple-payment-api/app/database"
)

func main() {
	app := app.New()
	app.DB = database.Init()

	http.HandleFunc("/", app.Router.ServeHTTP)
	log.Fatal(http.ListenAndServe(":7777", nil))
}