package factorymethod

import "testing"

func doConnect(factory DatabaseConnFactory) {
	conn := factory.CreateDatabaseConn()
	println(conn.Connect())
}

func Test_FactoryMethod_DatabaseConn(t *testing.T) {
	doConnect(&MySQLConnFactory{})
	doConnect(&PostgreSQLConnFactory{})
}
