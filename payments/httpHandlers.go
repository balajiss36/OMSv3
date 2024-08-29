package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	pb "github.com/balajiss36/common/api"
	models "github.com/balajiss36/common/models"
	"github.com/streadway/amqp"
)

type PaymentHTTPHandler struct {
	channel *amqp.Channel
}

func NewPaymentHTTPHandler(channel *amqp.Channel) *PaymentHTTPHandler {
	return &PaymentHTTPHandler{channel}
}

func (c *PaymentHTTPHandler) handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error reading webhook payload: %s", err)
	}

	var webhook models.RazorpayWebhook
	err = json.Unmarshal(payload, &webhook)
	if err != nil {
		log.Fatalf("Error parsing webhook payload: %s", err)
		return
	}

	log.Printf("Webhook received: %+v", webhook)
	if webhook.Event == "payment.captured" {
		log.Printf("Payment captured for order %v", webhook.Payload.Payment.Entity.ID)

		order := &pb.Order{
			Status:     "paid",
			OrderID:    webhook.Payload.Payment.Entity.ID,
			CustomerID: webhook.AccountID,
		}

		marshalledOrder, err := json.Marshal(order)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		err = c.channel.Publish("", "order.paid", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        marshalledOrder,
		})
		if err != nil {
			log.Fatalf("Error publishing order.paid message: %s", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
