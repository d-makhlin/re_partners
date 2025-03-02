package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCalculateHandlerHappyCase(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/calculate", CalculateHandler)

	requestBody := `{"numbers":[23,31,53],"target":500000}`
	req, _ := http.NewRequest("POST", "/calculate", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"23": 2, "31": 7, "53": 9429}`, w.Body.String())
}

func TestCalculateHandlerSetNegativeSizes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/calculate", CalculateHandler)

	requestBody := `{"numbers":[-1,100,200],"target":500000}`
	req, _ := http.NewRequest("POST", "/calculate", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "all sizes should be positive, -1 is not"}`, w.Body.String())
}

func TestCalculateHandlerNoAmount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/calculate", CalculateHandler)

	requestBody := `{"numbers":[100,200],"target":0}`
	req, _ := http.NewRequest("POST", "/calculate", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "The amount should be positive, 0 is not"}`, w.Body.String())
}

func TestCalculateHandlerNoPackSizes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/calculate", CalculateHandler)

	requestBody := `{"numbers":[],"target":100}`
	req, _ := http.NewRequest("POST", "/calculate", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "You should provide pack sizes"}`, w.Body.String())
}
