package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type CreateRequestDto struct {
	MerchantNumber string `json:"merchantNumber"`
	MerchantName string `json:"merchantName"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type CreateResponseDto struct {
	PaymentId string `json:"paymentId"`
}

type StatusResponseDto struct {
	PaymentId string `json:"paymentId"`
	Status string `json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type PaymentStatus struct {
	MerchantNumber string `json:"merchantNumber"`
	MerchantName string `json:"merchantName"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type StatusType string
const (
	Processing StatusType = "PROCESSING"
	Success StatusType = "SUCCESS"
	Failed   StatusType = "FAILED"
)

var payments map[string]PaymentStatus

func main() {
	payments = make(map[string]PaymentStatus)

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/payment/", createPayment).Methods("POST")
	myRouter.HandleFunc("/payment/{id}", paymentStatus)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createPayment")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var request CreateRequestDto
	json.Unmarshal(reqBody, &request)

	paymentId := uuid.NewString()
	payment := PaymentStatus{
		MerchantNumber: request.MerchantNumber,
		MerchantName:   request.MerchantName,
		Amount:         request.Amount,
		Currency:       request.Currency,
		Status:         string(Processing),
		Created:        now(),
		Updated:        now(),
	}

	payments[paymentId] = payment

	create := CreateResponseDto{
		PaymentId: paymentId,
	}
	json.NewEncoder(w).Encode(create)
}

func paymentStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: paymentStatus")

	vars := mux.Vars(r)
	key := vars["id"]

	var paymentStatus PaymentStatus
	paymentStatus = processPayment(key)

	status := StatusResponseDto{
		PaymentId: key,
		Status:    paymentStatus.Status,
		Created:   paymentStatus.Created,
		Updated:   paymentStatus.Updated,
	}

	json.NewEncoder(w).Encode(status)
}

func processPayment(paymentId string) PaymentStatus {
	var paymentStatus PaymentStatus
	paymentStatus = payments[paymentId]

	if paymentStatus.Status != string(Processing) {
		return paymentStatus
	}

	processed := now()

	if processed.After(paymentStatus.Created.Add(10 * time.Second)) {
		if rand.Intn(5) == 5 {
			paymentStatus.Status = string(Failed)
		} else {
			paymentStatus.Status = string(Success)
		}
		paymentStatus.Updated = now()
	}

	return paymentStatus
}

func now () time.Time {
	return time.Time.UTC(time.Now())
}