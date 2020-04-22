package controllers

import (
	"bytes"
	"net/http"

	"taxi/models"
	"taxi/queue"
)

// RequestController handles the random order from queue
func RequestController(w http.ResponseWriter, r *http.Request) {
	bookQ := &queue.Book{
		Order: make(chan models.Order),
	}

	queue.ProduceBook(bookQ)

	b := bytes.Buffer{}

	select {
	case order := <-bookQ.Order:
		b.WriteString(order.Name)
	}

	b.WriteTo(w)
}
