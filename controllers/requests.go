package controllers

import (
	"bytes"
	"fmt"
	"net/http"

	"taxi/models"
	"taxi/queue"
)

// RequestController handles orders from queue
func RequestsAdminController(w http.ResponseWriter, r *http.Request) {
	auditQ := &queue.Audit{
		Orders: make(chan []models.Order),
	}

	queue.ProduceReport(auditQ)

	b := bytes.Buffer{}

	select {
	case orders := <-auditQ.Orders:
		for _, order := range orders {
			if order.Booked > 0 {
				fmt.Fprintf(&b, "%s - %d\n", order.Name, order.Booked)
			}
		}
	}

	b.WriteTo(w)
}
