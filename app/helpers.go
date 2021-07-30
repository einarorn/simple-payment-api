package app

import (
	"math/rand"
	"simple-payment-api/app/models"
	"time"
)

func (app *App) processPayment(paymentId string) models.PaymentStatus {
	var paymentStatus models.PaymentStatus
	paymentStatus = app.DB.GetPayment(paymentId)

	if paymentStatus.Status != string(models.Processing) {
		return paymentStatus
	}

	processed := now()

	if processed.After(paymentStatus.Created.Add(10 * time.Second)) {
		if rand.Intn(5) == 5 {
			paymentStatus.Status = string(models.Failed)
		} else {
			paymentStatus.Status = string(models.Success)
		}
		paymentStatus.Updated = now()
	}

	return paymentStatus
}

func now() time.Time {
	return time.Time.UTC(time.Now())
}
