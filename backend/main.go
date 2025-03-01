package main

import (
	"backend/calculator"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var calc = calculator.NewCalculator([]int{})

func calculateHandler(ctx *gin.Context) {
	var data struct {
		Numbers []int `json:"numbers"`
		Target  int   `json:"target"`
	}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := calc.SetPackSizes(data.Numbers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := calc.CalculatePacks(data.Target)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())
	r.POST("/calculate", calculateHandler)

	r.Run(":8080")
}
