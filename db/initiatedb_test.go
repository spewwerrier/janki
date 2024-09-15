package db

import (
	"database/sql"
	"testing"
)

func ConnectTest(t *testing.T) {
	conn_str := "user=janki dbname=janki password=janki sslmode=disable port=5555"
	db := NewConnection(conn_str)
	err := db.CleanDb()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.CreateNewUser("dummyuser", "dummy password")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.CreateNewUser("dummyuser", "different password")
	if err != sql.ErrNoRows {
		t.Fatal("should complain about duplicate user but did not")
	}

	_, err = db.CreateNewUser("dummyuser2", "different password")
	if err != nil {
		t.Fatal(err)
	}

}
