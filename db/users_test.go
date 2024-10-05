package db

import (
	"testing"

	"janki/jlog"
)

func TestGetUsers(t *testing.T) {
	db := NewConnection("user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556", "/tmp/testfile.log")

	err := db.Create_db()
	defer db.raw.Close()
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.CreateNewUser("dummyuser", "dummy password", "groot", "groot")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.CreateNewUser("dummyuser", "different password", "groot", "groot")
	if err != jlog.ErrApiMultipleUsers {
		t.Fatal("should complain about duplicate user but did not")
	}

	_, err = db.CreateNewUser("dummyuser2", "different password", "groot", "groot")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDescriptions(t *testing.T) {
	db := NewConnection("user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556", "/tmp/testfile.log")

	err := db.Create_db()
	defer db.raw.Close()
	if err != nil {
		t.Fatal(err)
	}

	session_key, err := db.CreateNewUser("spw", "spewed everywhere", "bleh", "bleh")
	if err != nil {
		t.Fatal(err)
	}
	second_key, err := db.RetriveUserSession("spw", "spewed everywhere")
	if err != nil {
		t.Fatal(err)
	}
	if session_key != second_key {
		t.Fatal(err)
	}

	err = db.UpdateUser(session_key, "zahallo", "zahallo")
	if err != nil {
		t.Fatal(err)
	}

	third_key, err := db.RegenerateSessionKey("spw", "spewed everywhere")
	if err != nil {
		t.Fatal(err)
	}
	if third_key == second_key {
		t.Fatal(err)
	}
}
