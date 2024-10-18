package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"janki/api"
	"janki/db"
	"janki/jlog"
)

func Server() {
	database := db.NewConnection(os.Getenv("DB_URL"), "dblogs.log")
	defer database.Close()
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
		Addr:        "0.0.0.0:8080",
		Handler:     Middleware(Handler(api)),
		ReadTimeout: time.Second * 2,
		BaseContext: nil,
	}
	log.Println("listening on http://localhost:8080")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s.ListenAndServe()
	}()
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	fmt.Println("closing server")
}
