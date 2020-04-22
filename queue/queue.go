package queue

import (
	"time"

	"taxi/models"
	"taxi/utilrand"
)

const newOrdersInterval = 200 * time.Millisecond

// Audit holds the channel through which the orders will be returned
type Audit struct {
	Orders chan []models.Order
}

// Book holds the channel through which the random order will be returned
type Book struct {
	Order chan models.Order
}

var orders []models.Order

var booking = make(chan *Book)

var auditing = make(chan *Audit)

// ProduceBook sends a book request
func ProduceBook(book *Book) {
	booking <- book
}

// ProduceReport sends an audit request
func ProduceReport(audit *Audit) {
	auditing <- audit
}

// Consume listens and processes incoming book and audit requests
func Consume(done chan struct{}) {
	defer func() {
		close(booking)
		close(auditing)
	}()

	ticker := time.NewTicker(newOrdersInterval)
	defer ticker.Stop()

	for {
		select {
		case book := <-booking:
			n := utilrand.Num(len(orders))

			orders[n].Booked++

			book.Order <- orders[n]

			close(book.Order)

		case audit := <-auditing:
			cp := make([]models.Order, len(orders))

			copy(cp, orders)

			audit.Orders <- cp

			close(audit.Orders)

		case <-ticker.C:
			n := utilrand.Num(len(orders))

			orders = append(orders[:n], orders[n+1:]...)

			order := models.Order{
				Name: utilrand.Str(2),
			}

			orders = append(orders, order)

		case <-done:
			return
		}
	}
}
