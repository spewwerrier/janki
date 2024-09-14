package api

import (
	"janki/db"
	"net/http"
)

type Knob struct {
	Fields
	DB *db.Database
}

func (k Knob) Create(w http.ResponseWriter, r *http.Request) {

}
