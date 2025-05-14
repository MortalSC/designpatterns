package factorymethod

// DatabaseConn is an interface that defines the methods for database connections
type DatabaseConn interface {
	Connect() string
}

// MySQLConn is a concrete implementation of DatabaseConn for MySQL
type MySQLConn struct{}

func (m *MySQLConn) Connect() string {
	return "Connected to MySQL database"
}

// PostgreSQLConn is a concrete implementation of DatabaseConn for PostgreSQL
type PostgreSQLConn struct{}

func (p *PostgreSQLConn) Connect() string {
	return "Connected to PostgreSQL database"
}

// TODO: Add more database connection types as needed

// ----------------------------------------------------------------------------------------------------------------

// DatabaseConnFactory is an interface that defines the method for creating database connections
type DatabaseConnFactory interface {
	CreateDatabaseConn() DatabaseConn
}

// MySQLConnFactory is a concrete implementation of DatabaseConnFactory that creates MySQLConn
type MySQLConnFactory struct{}

func (m *MySQLConnFactory) CreateDatabaseConn() DatabaseConn {
	return &MySQLConn{}
}

// PostgreSQLConnFactory is a concrete implementation of DatabaseConnFactory that creates PostgreSQLConn
type PostgreSQLConnFactory struct{}

func (p *PostgreSQLConnFactory) CreateDatabaseConn() DatabaseConn {
	return &PostgreSQLConn{}
}

// TODO: Add more database connection factories as needed
