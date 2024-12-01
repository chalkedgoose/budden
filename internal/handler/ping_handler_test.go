package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	router := gin.Default()
	router.GET("/ping", Ping)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "pong"}`, w.Body.String())
}
