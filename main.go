package main

import (
	"log"

	"Book-API-Gin_Golang/handlers"
	"Book-API-Gin_Golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := models.NewStore()

	handlers.RegisterRoutes(r, store)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
