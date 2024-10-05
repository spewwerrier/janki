package jlog

import (
	"log"
	"net/http"
	"os"
	"time"
)

// https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

type Jlog struct {
	*log.Logger
	Filename string
}

func NewLogger(filename string) Jlog {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	// w := io.MultiWriter(os.Stdout, file)
	l := log.New(file, time.Now().Format(time.RFC1123)+" ----> ", 0)

	return Jlog{
		Logger:   l,
		Filename: filename,
	}
}

func (L *Jlog) WarningHttp(statuscode int, value string, w http.ResponseWriter) {
	L.Warning(value)
	w.WriteHeader(statuscode)
	_, _ = w.Write([]byte(value))
}

func (L *Jlog) ErrorHttp(statuscode int, value string, w http.ResponseWriter) {
	L.Error(value)
	w.WriteHeader(statuscode)
	_, _ = w.Write([]byte(value))
}

func (L *Jlog) InfoHttp(statuscode int, value string, w http.ResponseWriter) {
	L.Info(value)
	w.WriteHeader(statuscode)
	_, _ = w.Write([]byte(value))
}

func (L *Jlog) Warning(value string) {
	log.Printf(Yellow + "Warning: " + value + Reset)
	L.Println(value)
}

func (L *Jlog) Info(value string) {
	log.Printf(Green + "Info: " + value + Reset)
	L.Println(value)
}

func (L *Jlog) Error(value string) {
	log.Printf(Red + "Error: " + value + Reset)
	L.Println(value)
}
