package db

import (
	"database/sql"
	"sync"

	"janki/jlog"
)

type Database struct {
	log jlog.Jlog
	raw *sql.DB
	mu  *sync.Mutex
}
