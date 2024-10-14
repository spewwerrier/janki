package api

import (
	"encoding/json"
	"net/http"

	"janki/db"
	"janki/jlog"
)

type Knob struct {
	Fields
	DB  *db.Database
	Log jlog.Jlog
}

func (k Knob) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
		return
	}

	api_key := r.Form.Get("api_key")
	knob_name := r.Form.Get("knob_name")
	k.Log.Info(api_key)
	k.Log.Info(knob_name)

	knob := db.Knob{
		KnobName: knob_name,
		IsPublic: true,
	}

	err = k.DB.CreateNewKnob(api_key, knob)
	if err != nil {
		k.Log.Error(err.Error())
		k.Log.ErrorHttp(http.StatusInternalServerError, "cannot create new knob", w)
		return
	}
	k.Log.InfoHttp(http.StatusOK, "created new knob", w)
}

func (k Knob) Read(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
		return
	}
	api_key := r.Form.Get("api_key")
	knobs, err := k.DB.GetUserKnobs(api_key)
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "failed to get knob", w)
		return
	}
	knobs_json, err := json.Marshal(knobs)
	if err != nil {
		k.Log.ErrorHttp(http.StatusInternalServerError, "failed to encode the knobs", w)
		return
	}
	w.Write(knobs_json)
}

func (k Knob) UpdateKnobDescription(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
	// 	return
	// }

	// session_key := r.Form.Get("session_key")
	// knob_name := r.Form.Get("knob_name")
}
