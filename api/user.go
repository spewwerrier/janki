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
	session_key, err := u.DB.CreateNewUser(username, password)
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

func (u Users) Error(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("error on user"))
}
