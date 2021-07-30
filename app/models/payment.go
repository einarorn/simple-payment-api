package models

import "time"

type PaymentStatus struct {
	MerchantNumber string
	MerchantName   string
	Amount         string
	Currency       string
	Status         string
	Created        time.Time
	Updated        time.Time
}

type StatusType string
const (
	Processing StatusType = "PROCESSING"
	Success StatusType = "SUCCESS"
	Failed   StatusType = "FAILED"
)