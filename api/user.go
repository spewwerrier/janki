package api

import (
	"fmt"
	"janki/db"
	"log"
	"net/http"
)

type Users struct {
	Fields
	DB *db.Database
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	session_key, err := u.DB.CreateNewUser(username, password, "https://example.com", "I am groot")
	if err != nil {
		_, _ = w.Write([]byte("duplicate user"))
		fmt.Println(err)
	}
	cookie := http.Cookie{
		Name:   "user_token",
		Value:  session_key,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(w, &cookie)
	_, _ = w.Write([]byte(cookie.Value))
}

func (u Users) CreateDescription(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	session_key := r.Form.Get("session_key")
	image_url := r.Form.Get("image_url")
	description := r.Form.Get("description")
	err = u.DB.CreateUserDescription(session_key, image_url, description)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte("cannot create again"))
		return
	}
	_, _ = w.Write([]byte("created description"))
}

func (u Users) UpdateUserDescription(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	session_key := r.Form.Get("session_key")
	image_url := r.Form.Get("image_url")
	description := r.Form.Get("description")
	err = u.DB.UpdateUser(session_key, image_url, description)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte("cannot update description"))
		return
	}
	_, _ = w.Write([]byte("updated description"))
}

func (u Users) Read(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	session_key, err := u.DB.RetriveUserSession(username, password)
	if err != nil {
		_, _ = w.Write([]byte("cannot retrive user id from the credentials"))
		return
	}
	fmt.Println(session_key)
	_, _ = w.Write([]byte(session_key))

}

func (u Users) Error(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("error on user"))
}
