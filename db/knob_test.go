package db

import (
	"errors"
	"testing"

	"janki/jlog"
)

func TestKnob(t *testing.T) {
	db := NewConnection("user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556", "/tmp/testfile.log")

	err := db.Create_db()
	defer db.raw.Close()
	if err != nil {
		t.Fatal(err)
	}

	user1, err := db.CreateNewUser("aagaman", "hello", "", "")
	if err != nil {
		t.Fatal(err)
	}

	user2, err := db.CreateNewUser("aagaman2", "hello", "", "")
	if err != nil {
		t.Fatal(err)
	}

	k := Knob{
		KnobName: "writing a protocol in c",
		IsPublic: true,
	}
	err = KnobCreate(db, user1, k)
	if err != nil {
		t.Fatal(err)
	}

	err = KnobCreate(db, user2, k)
	if err != nil {
		t.Fatal(err)
	}

	err = KnobCreate(db, user1, k)
	if err != jlog.ErrKnobAlreadyExists {
		t.Fatal(errors.New("should complain about multiple knobs but did not"))
	}
}

func KnobCreate(db *Database, user string, k Knob) error {
	err := db.CreateNewKnob(user, k)
	if err != nil {
		return err
	}

	err = db.CreateNewKnob(user, k)
	if err == nil {
		return errors.New("supposed to return error on duplicate creation but did not")
	}
	return nil
}

func TestReads(t *testing.T) {
	db := NewConnection("user=janki_test dbname=janki_test password=janki_test sslmode=disable port=5556", "/tmp/testfile.log")

	err := db.Create_db()
	defer db.raw.Close()
	if err != nil {
		t.Fatal(err)
	}

	user1, err := db.CreateNewUser("aagaman", "hello", "", "")
	if err != nil {
		t.Fatal(err)
	}
	k := Knob{
		KnobName: "something",
		IsPublic: true,
	}
	KnobCreate(db, user1, k)
}
