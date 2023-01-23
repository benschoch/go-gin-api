package ingredient

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define the root endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Create a new request to the root endpoint
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert that the response has a status code of 200
	if w.Code != 200 {
		t.Errorf("Unexpected status code: %d", w.Code)
	}

	// Assert that the response body contains the string "Hello, World!"
	if !strings.Contains(w.Body.String(), "Hello, World!") {
		t.Errorf("Unexpected response body: %s", w.Body.String())
	}
}
