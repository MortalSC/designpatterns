package functionaloptions

import "time"

// DBClient is a struct that represents a database client with various options.
type DBClient struct {
	Address string
	Timeout time.Duration
	Retries time.Duration
	UseSSL  bool
}

// Option is type for a function that modifies the DBClient.
type DBOption func(*DBClient)

func WithAddress(addr string) DBOption {
	return func(c *DBClient) {
		c.Address = addr
	}
}

func DBWithTimeout(timeout time.Duration) DBOption {
	return func(c *DBClient) {
		c.Timeout = timeout
	}
}

func WithRetries(retries time.Duration) DBOption {
	return func(c *DBClient) {
		c.Retries = retries
	}
}

func WithSSL(enable bool) DBOption {
	return func(c *DBClient) {
		c.UseSSL = enable
	}
}

func NewDBClient(options ...DBOption) *DBClient {
	// Default options
	client := &DBClient{
		Address: "localhost:3306",
		Timeout: 10 * time.Second,
		Retries: 3 * time.Second,
		UseSSL:  false,
	}

	// Apply options to the client
	for _, option := range options {
		option(client)
	}

	return client
}
