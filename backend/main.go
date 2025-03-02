package main

import (
	"github.com/gin-contrib/cors"

	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())
	r.POST("/calculate", handlers.CalculateHandler)

	r.Run(":8080")
}
