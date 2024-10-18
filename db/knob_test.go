package db

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestKnob(t *testing.T) {
	db := NewConnection("user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556", "/tmp/testfile.log")

	err := db.Create_db()
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	tx, err := db.raw.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	rows, err := tx.Query("select * from knobs where id = 1")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		fmt.Println(rows)
	}
}

func Execute(db *Database, query string, args ...interface{}) (sql.Result, error) {
	for _, arg := range args {
		fmt.Println(arg)
	}
	return nil, nil
}
