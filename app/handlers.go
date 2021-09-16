package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"simple-payment-api/app/models"
	"time"
)

type LoginRequestDto struct {
	Username string `json:"username"`
	Password   string `json:"password"`
}

type LoginResponseDto struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int `json:"expires_in"`
}

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

func (app *App) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Login] called...")

		reqBody, _ := ioutil.ReadAll(r.Body)
		var request LoginRequestDto
		json.Unmarshal(reqBody, &request)

		if request.Username == "eoo" && request.Password == "Cowabunga!" {
			token, _ := getToken(request.Username)

			response := LoginResponseDto{
				AccessToken: token,
				TokenType: "bearer",
				ExpiresIn: ValidMinutes * 60,
			}
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Name and password do not match"))
		}
	}
}

func (app *App) CreatePaymentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[CreatePayment] called...")

		reqBody, _ := ioutil.ReadAll(r.Body)
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

		json.NewEncoder(w).Encode(create)
	}
}

func (app *App) PaymentStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[PaymentStatus] called...")

		vars := mux.Vars(r)
		key := vars["id"]

		var paymentStatus models.PaymentStatus
		paymentStatus = app.processPayment(key)

		status := StatusResponseDto{
			PaymentId: key,
			Status:    paymentStatus.Status,
			Created:   paymentStatus.Created,
			Updated:   paymentStatus.Updated,
		}

		json.NewEncoder(w).Encode(status)
	}
}