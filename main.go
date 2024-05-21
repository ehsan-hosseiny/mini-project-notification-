package main

import (
	"notification/entities"
	"notification/services"
)

func main() {
	order := entities.Order{
		ID:           1,
		UserFullName: "john Doe",
		UserId:       "9121234567",
		Price:        float64(100),
		Status:       true,
		NotificationType: entities.Email,
	}

	orderService := services.NewOrderService()
	orderService.CreateOrder(&order)

}
