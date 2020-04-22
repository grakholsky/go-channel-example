package queue

import (
	"taxi/models"
	"taxi/utilrand"
)

func Seed() {
	for i := 0; i < 50; i++ {
		order := models.Order{
			Name: utilrand.Str(2),
		}

		orders = append(orders, order)
	}
}
