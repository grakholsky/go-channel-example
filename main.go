package main

import (
	"log"
	"net/http"

	"taxi/queue"

	c "taxi/controllers"
)

func main() {
	done := make(chan struct{})
	defer func() {
		done <- struct{}{}
	}()

	http.HandleFunc("/request", c.RequestController)
	http.HandleFunc("/admin/requests", c.RequestsAdminController)

	queue.Seed()
	go queue.Consume(done)

	log.Printf("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
