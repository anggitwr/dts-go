package main

import (
	"versi2/controller"
	"versi2/lib"
	"versi2/model"
	"versi2/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Person{}, model.Product{})
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personController := controller.NewPersonController(personRepository)

	ginEngine := gin.Default()

	ginEngine.GET("/person", personController.GetAll)
	ginEngine.POST("person", personController.Create)

	err = ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}
