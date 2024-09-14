package db

import (
	"testing"
)

func ConnectTest(t *testing.T) {
	conn_str := "user=janki dbname=janki password=janki sslmode=disable port=5555"
	TestDB := NewConnection(conn_str)
	_ = TestDB.GetUser()
}
