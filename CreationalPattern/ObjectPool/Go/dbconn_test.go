package objectpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DBConnectionPool(t *testing.T) {
	assert := assert.New(t)

	pool := NewDBConnectionPool(3)

	conn1, err := pool.Acquire(1 * time.Second)
	assert.NoError(err)
	assert.Equal(1, conn1.ID)

	conn2, err := pool.Acquire(1 * time.Second)
	assert.NoError(err)
	assert.Equal(2, conn2.ID)

	conn3, err := pool.Acquire(1 * time.Second)
	assert.NoError(err)
	assert.Equal(3, conn3.ID)

	// Test timeout
	conn4, err := pool.Acquire(1 * time.Second)
	assert.NoError(err)
	assert.Equal(4, conn4.ID)

	// Release connections back to the pool
	pool.Release(conn1)
	pool.Release(conn2)
	pool.Release(conn3)
	pool.Release(conn4)
}
