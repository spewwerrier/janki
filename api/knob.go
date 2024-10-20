package api

import (
	"encoding/json"
	"fmt"
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
	// this gives identifier
	identifier := r.Form.Get("fork")

	k.Log.Info(api_key)
	k.Log.Info(knob_name)
	k.Log.Info(identifier)

	if len(identifier) > 1 {
		err := k.DB.ForkKnob(api_key, identifier)
		if err != nil {
			k.Log.ErrorHttp(http.StatusInternalServerError, "cannot fork knob", w)
		}
		return
	}

	if len(knob_name) < 1 {
		k.Log.ErrorHttp(http.StatusInternalServerError, "cannot create new knob, too short name", w)
		return
	}

	knob := db.KnobDescription{
		Knob: db.Knob{
			KnobName: knob_name,
			IsPublic: true,
		},
		Description: "this is a dummy description",
	}

	knob_identifier, err := k.DB.CreateNewKnob(api_key, knob)
	if err != nil {
		k.Log.Error(err.Error())
		k.Log.ErrorHttp(http.StatusInternalServerError, "cannot create new knob", w)
		return
	}
	k.Log.InfoHttp(http.StatusOK, knob_identifier, w)
}

// if read has a parameter -knobid = something then it returns that something
// else it returns everything
func (k Knob) Read(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
		return
	}
	api_key := r.Form.Get("api_key")
	identifier := r.Form.Get("knob_id")

	if identifier != "" {
		knob, _ := k.DB.GetKnobDescriptions(api_key, identifier)
		knobs_json, err := json.Marshal(knob)
		if err != nil {
			k.Log.ErrorHttp(http.StatusInternalServerError, "failed to encode the knobs", w)
			return
		}
		fmt.Println("reading knob", knob)
		w.Write(knobs_json)
		return
	}

	knobs, err := k.DB.GetUserKnobs(api_key)
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "failed to get knob"+err.Error(), w)
		return
	}
	knobs_json, err := json.Marshal(knobs)
	if err != nil {
		k.Log.ErrorHttp(http.StatusInternalServerError, "failed to encode the knobs ", w)
		return
	}
	fmt.Println(knobs)
	w.Write(knobs_json)
}

// update a single knob
func (k Knob) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
		return
	}

	api_key := r.Form.Get("api_key")
	knob_id := r.Form.Get("knob_id")

	ispublic := r.Form.Get("ispublic")
	ques := r.Form.Get("questions")
	refs := r.Form.Get("refs")

	if refs != "" {
		fmt.Println(api_key, knob_id, ques)
		k.DB.UpdateKnobDescriptions(api_key, knob_id, "refs", refs)
	}
	if ques != "" {
		fmt.Println("wtmoooo", api_key, knob_id, refs)
		k.DB.UpdateKnobDescriptions(api_key, knob_id, "ques", ques)
	}
	if ispublic != "" {
		fmt.Println("wtmoooo", api_key, knob_id, ispublic)
		k.DB.UpdateKnobPublic(api_key, knob_id, ispublic)

	}
}
