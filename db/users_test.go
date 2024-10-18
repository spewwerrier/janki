package db

import (
	"testing"

	"janki/jlog"
)

func ConnectTestInstance(t *testing.T) *Database {
	conn_str := "postgres://janki_test:janki_test@localhost/janki_test?sslmode=disable&port=5556"
	db := NewConnection(conn_str, "/tmp/testfile.log")

	err := db.Create_db()
	if err != nil {
		t.Fatal(err)
	}
	return db
}

const (
	username    = "dummyuser"
	passwd      = "dummypass"
	description = "hello I am a dummy dum dum"
	image_url   = "https://example.com/image.png"
)

func TestUsers(t *testing.T) {
	d := ConnectTestInstance(t)

	api, err := d.CreateNewUser(username, passwd)
	if err != nil {
		t.Fatal(err)
		return
	}

	err = d.UpdateUser(api, "description", description)
	if err != nil {
		t.Fatal(err)
		return
	}
	d.UpdateUser(api, "image_url", image_url)
	if err != nil {
		t.Fatal(err)
		return
	}

	api_check, err := d.RetriveUserApi(username, passwd)
	if err != nil || api_check != api {
		t.Fatal(err)
		return
	}

	user, err := d.RetriveUser(api)
	if err != nil {
		t.Fatal(err)
		return
	}

	if user.User.Name != username {
		t.Fatalf("Username is not correct")
		return
	}

	if user.Description != description {
		t.Fatalf("Description is not correct")
		return
	}

	if user.Image_url != image_url {
		t.Fatalf("Image url is not correct")
		return
	}

	_, err = d.CreateNewUser(username, passwd)
	// err if user tries to create another user with same name
	if err != jlog.ErrApiMultipleUsers {
		t.Fatal("should throw multiple user exists error but did not")
		return
	}
}
