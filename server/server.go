package server

import (
	"janki/api"
	"janki/db"
	"net/http"
)

func Server() {
	database := db.NewConnection("user=janki dbname=janki password=janki sslmode=disable port=5555")

	api := &api.Api{
		Users: api.Users{DB: database},
		Knob:  api.Knob{DB: database},
	}

	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: Handler(api),
	}
	_ = s.ListenAndServe()
}
