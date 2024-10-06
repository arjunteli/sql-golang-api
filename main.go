package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Define the route and associate it with the controller function
	r.GET("/person/:person_id/info", GetPersonInfo)
	r.POST("/person/create", CreatePerson)

	// server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
