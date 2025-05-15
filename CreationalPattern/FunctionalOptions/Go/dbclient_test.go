package functionaloptions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DBClient(t *testing.T) {
	assert := assert.New(t)

	// Test the DBClient with different options
	client := NewDBClient(
		WithAddress("localhost:5432"),
		DBWithTimeout(5*time.Second),
		WithRetries(2*time.Second),
		WithSSL(true),
	)

	assert.Equal(client.Address, "localhost:5432")
	assert.Equal(client.Timeout, 5*time.Second)
	assert.Equal(client.Retries, 2*time.Second)
	assert.Equal(client.UseSSL, true)
}
