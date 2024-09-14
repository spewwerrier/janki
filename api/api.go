package api

import (
	"net/http"
)

type Fields interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Error(w http.ResponseWriter, r *http.Request)
}

type Api struct {
	Users Users
	Knob  Knob
}
