package chainofresponsibility

import (
	"net/http"
	"testing"
)

func Test_HTTPRequest(t *testing.T) {
	// Create the chain of handlers
	logger := &LoggerHandler{}
	auth := &AuthHandler{}
	business := &BusinessHandler{}

	// Set the next handler in the chain
	logger.SetNext(auth)
	auth.SetNext(business)

	// Create a sample HTTP request
	request, _ := http.NewRequest("GET", "http://example.com", nil)

	// Process the request through the chain
	logger.Handle(request)
}
