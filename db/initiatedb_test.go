package db

import (
	"testing"
)

func TestConnect(t *testing.T) {
	conn_str := "user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556"
	db := NewConnection(conn_str)
	err := db.CleanDb()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Create_db()
	if err != nil {
		t.Fatal(err)
	}
}
