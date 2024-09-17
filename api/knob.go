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
	err := r.ParseForm()
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
		return
	}
	session_key := r.Form.Get("session_key")
	knob_name := r.Form.Get("knob_name")
	err = k.DB.CreateNewKnob(session_key, knob_name)
	if err != nil {
		k.Log.Error(err.Error())
		k.Log.ErrorHttp(http.StatusInternalServerError, "cannot create new knob", w)
		return
	}
	k.Log.InfoHttp(http.StatusOK, "created new knob", w)
}
