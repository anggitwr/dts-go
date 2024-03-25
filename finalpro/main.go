package main

import (
	"finalpro/lib"
	"finalpro/router"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.StartDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	router.StartRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
