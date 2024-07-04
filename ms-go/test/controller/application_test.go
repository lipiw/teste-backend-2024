package controller

import (
	"encoding/json"
	"log"
	"ms-go/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndexHome(t *testing.T) {
	expectedBody := gin.H{
		"message": "[ms-go] | Success",
		"status":  http.StatusOK,
	}

	r := router.SetupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)

	log.Printf("Expected message: %s", expectedBody["message"])
	log.Printf("Response message: %s", response["message"])

	assert.Equal(t, expectedBody["message"], response["message"])
	assert.Equal(t, expectedBody["status"], int(response["status"].(float64)))
}
