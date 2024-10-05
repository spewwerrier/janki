package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"janki/api"
	"janki/db"
	"janki/jlog"
)

func Server() {
	database := db.NewConnection("user=janki dbname=janki password=janki sslmode=disable port=5555", "dblogs.log")
	logger := jlog.NewLogger("apilogs.log")
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
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			s.Close()
			panic(err)
		}
	}()
	<-c
	s.Close()
	fmt.Println("closing server")
}
