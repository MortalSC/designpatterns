package chainofresponsibility

import (
	"net/http"
)

// The following is an example of an HTTP request processing chain,
// which includes three steps: logging, authentication, and business processing

// ---------- Processor interface ----------
type Handler interface {
	Handle(request *http.Request)
	SetNext(handler Handler)
}

// ---------- Base processor (implementing chain transfer logic) ----------
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) invokeNext(request *http.Request) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// ---------- Specific processor: logging ----------

type LoggerHandler struct {
	BaseHandler
}

func (l *LoggerHandler) Handle(request *http.Request) {
	// Log the request
	println("Logging request:", request.URL.String())
	// Invoke the next handler in the chain
	l.invokeNext(request)
}

// ---------- Specific processor: authentication ----------
type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(request *http.Request) {
	// Authenticate the request
	println("Authenticating request:", request.URL.String())
	// Invoke the next handler in the chain
	a.invokeNext(request)
}

// ---------- Specific processor: business processing ----------
type BusinessHandler struct {
	BaseHandler
}

func (b *BusinessHandler) Handle(request *http.Request) {
	// Process the business logic
	println("Processing business logic for request:", request.URL.String())
	// Invoke the next handler in the chain
	// b.invokeNext(request)
	// No next handler, so no further processing
}

// ---------- Example usage ----------
// func ExampleUsage() {
// 	// Create the chain of handlers
// 	logger := &LoggerHandler{}
// 	auth := &AuthHandler{}
// 	business := &BusinessHandler{}

// 	// Set up the chain
// 	logger.SetNext(auth)
// 	auth.SetNext(business)

// 	validRequest := &http.Request{
// 		Method: "GET",
// 		URL:    &http.URL{Path: "/api/data"},
// 		Header: map[string][]string{"Authorization": {"valid-token"}},
// 	}
// 	invalidRequest := &http.Request{
// 		Method: "POST",
// 		URL:    &http.URL{Path: "/api/admin"},
// 		Header: map[string][]string{"Authorization": {"invalid-token"}},
// 	}

// 	fmt.Println("---- 有效请求处理 ----")
// 	logger.Handle(validRequest)

// 	fmt.Println("\n---- 无效请求处理 ----")
// 	logger.Handle(invalidRequest)
// }
