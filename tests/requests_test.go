package tests

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppEndpoints(t *testing.T) {
	// Create a new HTTP request against your test router
	req, err := http.NewRequest("GET", "/heroes", nil)
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request using the test router
	Router.ServeHTTP(rr, req)

	// Assert the response status code and body
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "")
}
