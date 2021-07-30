package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"simple-payment-api/app/models"
	"time"
)

type CreateRequestDto struct {
	MerchantNumber string `json:"merchantNumber"`
	MerchantName   string `json:"merchantName"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
}

type CreateResponseDto struct {
	PaymentId string `json:"paymentId"`
}

type StatusResponseDto struct {
	PaymentId string    `json:"paymentId"`
	Status    string    `json:"status"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

func (app *App) CreatePaymentHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Endpoint hit: createPayment")

		reqBody, _ := ioutil.ReadAll(req.Body)
		var request CreateRequestDto
		json.Unmarshal(reqBody, &request)

		payment := models.PaymentStatus{
			MerchantNumber: request.MerchantNumber,
			MerchantName:   request.MerchantName,
			Amount:         request.Amount,
			Currency:       request.Currency,
			Status:         string(models.Processing),
			Created:        now(),
			Updated:        now(),
		}

		paymentId := app.DB.CreatePayment(payment)

		create := CreateResponseDto{
			PaymentId: paymentId,
		}
		json.NewEncoder(writer).Encode(create)
	}
}

func (app *App) PaymentStatusHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		fmt.Println("Endpoint hit: paymentStatus")

		vars := mux.Vars(req)
		key := vars["id"]

		var paymentStatus models.PaymentStatus
		paymentStatus = app.processPayment(key)

		status := StatusResponseDto{
			PaymentId: key,
			Status:    paymentStatus.Status,
			Created:   paymentStatus.Created,
			Updated:   paymentStatus.Updated,
		}

		json.NewEncoder(writer).Encode(status)
	}
}