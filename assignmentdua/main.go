package main

import (
	"assignmentdua/controller"
	"assignmentdua/lib"
	"assignmentdua/model"
	"assignmentdua/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Order{}, model.Item{})
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	ginEngine := gin.Default()
	ginEngine.GET("/orders", orderController.GetAll)
	ginEngine.POST("order", orderController.Create)
	ginEngine.DELETE("orders/:id", orderController.Delete)
	ginEngine.PUT("orders/:id", orderController.UpdateOrder)
	err = ginEngine.Run("localhost: 8083")
	if err != nil {
		panic(err)
	}
}
