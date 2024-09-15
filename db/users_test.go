package db

import (
	"fmt"
	"testing"
)

func TestGetUsers(t *testing.T) {
	database := NewConnection("user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556")

	err := database.Create_db()
	if err != nil {
		t.Fatal(err)
	}

	s, err := database.RetriveUser("someradnomcookiestring")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("t: %v\n", s)
	fmt.Println(s)
}
