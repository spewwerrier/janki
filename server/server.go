package server

import (
	"janki/api"
	"janki/db"
	jankilog "janki/logs"
	"log"
	"net/http"
)

func Server() {
	database := db.NewConnection("user=janki dbname=janki password=janki sslmode=disable port=5555", "dblogs.log")
	logger := jankilog.NewLogger("apilogs.log")
	err := database.Create_db()
	if err != nil {
		log.Panic(err)
	}

	api := &api.Api{
		Users: api.Users{DB: database, Log: logger},
		Knob:  api.Knob{DB: database, Log: logger},
	}

	s := &http.Server{
		Addr:    "localhost:8080",
		Handler: Handler(api),
	}
	log.Println("listening on http://localhost:8080")
	err = s.ListenAndServe()
	if err != nil {
		log.Panic(err)
		s.Close()
	}
}
