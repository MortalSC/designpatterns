package functionaloptions

import (
	"net/http"
	"time"
)

// HTTPClient is a struct that represents an HTTP client with various options.
type HTTPClient struct {
	Timeout time.Duration
}

// DefaultHTTPClient is the default options for the HTTP client.
func DefaultHTTPClient() HTTPClient {
	return HTTPClient{
		Timeout: 10 * time.Second, // Default timeout: 10s
	}
}

// Option is a function that takes a pointer to HTTPClient and modifies it.
type Option func(*HTTPClient)

// WithTimeout is an option that sets the timeout for the HTTP client.
func WithTimeout(timeout time.Duration) Option {
	return func(hc *HTTPClient) {
		hc.Timeout = timeout
	}
}

func NewHTTPClient(options ...Option) *http.Client {
	client := DefaultHTTPClient()
	for _, option := range options {
		option(&client)
	}
	return &http.Client{
		Timeout: client.Timeout,
	}
}
