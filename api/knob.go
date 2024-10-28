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
	description := r.Form.Get("description")
	// if fork is mentioned then it should be with an identifier to which knob to fork
	// note that you cannot fork a private knob
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

	knob := db.Knob{
		KnobName:    knob_name,
		IsPublic:    true,
		Description: description,
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
	identifier := r.Form.Get("identifier")

	if identifier != "" {
		knob, _ := k.DB.RetrieveKnobDescriptions(api_key, identifier)

		knobs_json, err := json.Marshal(knob)
		if err != nil {
			k.Log.ErrorHttp(http.StatusInternalServerError, "failed to encode the knobs", w)
			return
		}
		fmt.Println("reading knob", knob)
		w.Write(knobs_json)
		return
	}

	knobs, err := k.DB.RetrieveUserKnobs(api_key)
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
	knob_id := r.Form.Get("identifier")
	ispublic := r.Form.Get("ispublic")

	v := map[string]string{
		"knobtopics":       r.Form.Get("topics"),
		"knobtodo":         r.Form.Get("todo"),
		"knobthingstoread": r.Form.Get("tor"),
		"knobreferences":   r.Form.Get("refs"),
		"knoburls":         r.Form.Get("urls"),
		"knobquestions":    r.Form.Get("ques"),
		"knobsuggestions":  r.Form.Get("suggestions"),
	}
	for table, content := range v {
		if content != "" {
			err = k.DB.InsertKnob(table, content, knob_id, api_key)
			if err != nil {
				fmt.Println(err)
				k.Log.ErrorHttp(http.StatusInternalServerError, "failed to update knob property "+content, w)
				return

			}
		}
	}

	// if ispublic != "" {
	// 	fmt.Println("wtmoooo", api_key, knob_id, ispublic)
	// 	k.DB.UpdateKnobPublic(api_key, knob_id, ispublic)
	// 	if err != nil {
	// 		k.Log.ErrorHttp(http.StatusInternalServerError, "failed to change knob visibility", w)
	// 		return

	// 	}
	// }
}

func (k Knob) Delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		k.Log.ErrorHttp(http.StatusBadRequest, "unable to parse form", w)
		return
	}

	api_key := r.Form.Get("api_key")
	knob_id := r.Form.Get("knob_id")

	k.DB.DeleteKnob(api_key, knob_id)
}
