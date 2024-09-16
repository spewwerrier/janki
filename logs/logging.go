package jankilog

import (
	"log"
	"net/http"
	"os"
	"time"
)

// https://www.dolthub.com/blog/2024-02-23-colors-in-golang/
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type JankiLog struct {
	Filename string
	*log.Logger
}

func NewLogger(filename string) JankiLog {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// w := io.MultiWriter(os.Stdout, file)
	l := log.New(file, time.Now().Format(time.RFC1123)+" ----> ", 0)

	return JankiLog{
		Logger:   l,
		Filename: filename,
	}
}

func (L *JankiLog) WarningHttp(statuscode int, value string, w http.ResponseWriter) {
	L.Warning(value)
	w.WriteHeader(statuscode)
	_, _ = w.Write([]byte(value))
}

func (L *JankiLog) ErrorHttp(statuscode int, value string, w http.ResponseWriter) {
	L.Error(value)
	w.WriteHeader(statuscode)
	_, _ = w.Write([]byte(value))
}

func (L *JankiLog) Warning(value string) {
	log.Printf(Yellow + "Warning: " + value + Reset)
	L.Println(value)
}

func (L *JankiLog) Info(value string) {
	log.Printf(Green + "Info: " + value + Reset)
	L.Println(value)
}

func (L *JankiLog) Error(value string) {
	log.Printf(Red + "Error: " + value + Reset)
	L.Println(value)
}
