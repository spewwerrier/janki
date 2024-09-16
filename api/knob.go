package api

import (
	"janki/db"
	jankilog "janki/logs"
	"net/http"
)

type Knob struct {
	Fields
	DB  *db.Database
	Log jankilog.JankiLog
}

func (k Knob) Create(w http.ResponseWriter, r *http.Request) {

}
