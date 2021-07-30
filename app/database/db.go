package database

import (
	"github.com/google/uuid"
	"simple-payment-api/app/models"
)

type DB struct {
	database map[string]models.PaymentStatus
}

func Init() *DB {
	db := &DB {
		database: make(map[string]models.PaymentStatus),
	}

	return db
}

func (db *DB) CreatePayment(payment models.PaymentStatus) string {
	paymentId := uuid.NewString()
	db.database[paymentId]	= payment

	return paymentId
}

func (db *DB) GetPayment(paymentId string) models.PaymentStatus {
	return db.database[paymentId]
}
