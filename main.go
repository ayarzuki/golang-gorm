package main

import (
	"fmt"
	"golang-gorm/configs"
	"golang-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.StartDB()
	if db == nil {
		fmt.Println("db failed to run!")
	}

	r := gin.Default()
	orderController := controllers.NewControllerOrder(db)
	userRoute := r.Group("/orders")
	{
		userRoute.GET("/", orderController.GetOrders)
		userRoute.POST("/", orderController.CreateOrder)
		userRoute.PUT("/:orderId", orderController.UpdateOrderByID)
		userRoute.DELETE("/:orderId", orderController.DeleteOrderByID)
	}

	r.Run(":8080")
}
