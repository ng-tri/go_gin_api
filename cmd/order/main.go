package main

import (
	"fmt"
	"go-gin-api/internal/order"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	orderService := order.NewOrderService()
	r.POST("/order/create", orderService.CreateOrder)

	fmt.Println("Order Service is running at :8082")
	r.Run(":8082") // Order Service on port 8082
}
