package handlers

import (
	"backend/calculator"
	"net/http"

	"github.com/gin-gonic/gin"
)

var calc = calculator.NewCalculator([]int{})

func CalculateHandler(ctx *gin.Context) {
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
