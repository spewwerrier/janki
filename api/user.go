package api

import (
	"janki/db"
	"net/http"
)

type Users struct {
	Fields
	DB *db.Database
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello"))
}

func (u Users) Error(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("error on user"))
}
