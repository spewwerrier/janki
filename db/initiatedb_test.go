package db

import (
	"testing"
)

func TestConnect(t *testing.T) {
	conn_str := "postgres://janki_test:janki_test@localhost/janki_test?sslmode=disable&port=5556"
	db := NewConnection(conn_str, "/tmp/testfile.log")
	defer db.Close()
	err := db.CleanDb()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Create_db()
	if err != nil {
		t.Fatal(err)
	}
}
