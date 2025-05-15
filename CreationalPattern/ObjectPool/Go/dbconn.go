package objectpool

import (
	"fmt"
	"sync"
	"time"
)

// DBconnection represents a database connection object.
type DBConnection struct {
	ID int
}

func (c *DBConnection) Query() string {
	return fmt.Sprintf("Querying database with connection ID: %d", c.ID)
}

// DBConnectionPool is a struct that represents a pool of database connections.
type DBConnectionPool struct {
	pool    chan *DBConnection
	maxSize int
	mu      sync.Mutex
	nextID  int
}

func NewDBConnectionPool(maxSize int) *DBConnectionPool {
	p := &DBConnectionPool{
		pool:    make(chan *DBConnection, maxSize),
		maxSize: maxSize,
		nextID:  1,
	}

	// Pre-fill the pool with database connections.
	for i := 0; i < 2; i++ {
		p.pool <- p.createConnection()
	}

	return p
}

func (p *DBConnectionPool) createConnection() *DBConnection {
	p.mu.Lock()
	defer p.mu.Unlock()

	fmt.Printf("Creating new DB connection with ID: %d\n", p.nextID)
	time.Sleep(500 * time.Millisecond) // Simulate time taken to create a connection
	conn := &DBConnection{ID: p.nextID}
	p.nextID++

	return conn
}

func (p *DBConnectionPool) Acquire(timeout time.Duration) (*DBConnection, error) {
	select {
	case conn := <-p.pool:
		return conn, nil
	case <-time.After(timeout): // support timeout
		if len(p.pool) < p.maxSize {
			conn := p.createConnection()
			return conn, nil
		}
		return nil, fmt.Errorf("timeout while acquiring DB connection")
	}
}

func (p *DBConnectionPool) Release(conn *DBConnection) {
	select {
	case p.pool <- conn:
	default:
		fmt.Printf("DB connection pool is full, closing connection with ID: %d\n", conn.ID)
	}
}
