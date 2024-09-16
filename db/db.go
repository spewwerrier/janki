package db

import (
	"database/sql"
	jankilog "janki/logs"
	"sync"
)

type Database struct {
	mu  sync.Mutex
	db  *sql.DB
	log jankilog.JankiLog
}
