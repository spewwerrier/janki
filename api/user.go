package api

import (
	"encoding/json"
	"net/http"

	"janki/db"
	"janki/jlog"
)

type Users struct {
	Fields
	DB  *db.Database
	Log jlog.Jlog
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}

	// this is just a simple validation, move this to new file with all input validations and such
	username := r.Form.Get("username")
	if len(username) < 3 {
		u.Log.ErrorHttp(http.StatusBadRequest, "username should be more than 3 characters long", w)
		return
	}

	password := r.Form.Get("password")
	if len(password) < 8 {
		u.Log.ErrorHttp(http.StatusBadRequest, "password should be more than 8 characters long", w)
		return
	}

	api_key, err := u.DB.CreateNewUser(username, password)
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "duplicate user", w)
		return
	}
	cookie := http.Cookie{
		Name:   "user_token",
		Value:  api_key,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(w, &cookie)
	_, _ = w.Write([]byte(cookie.Value))
}

func (u Users) CreateDescription(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}
	session_key := r.Form.Get("session_key")
	image_url := r.Form.Get("image_url")
	description := r.Form.Get("description")
	err = u.DB.CreateUserDescription(session_key, image_url, description)
	if err != nil {
		u.Log.ErrorHttp(http.StatusInternalServerError, "cannot create user description", w)
		return
	}
	_, _ = w.Write([]byte("created description"))
}

func (u Users) Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}
	api_key := r.Form.Get("api_key")

	description := r.Form.Get("description")
	image_url := r.Form.Get("iamge_url")

	if description != "" {
		err = u.DB.UpdateUser(api_key, "description", description)
		if err != nil {
			u.Log.ErrorHttp(http.StatusInternalServerError, "cannot update user description", w)
			return
		}
	}
	if image_url != "" {
		err = u.DB.UpdateUser(api_key, "image_url", image_url)
		if err != nil {
			u.Log.ErrorHttp(http.StatusInternalServerError, "cannot update user image", w)
			return
		}
	}

	_, _ = w.Write([]byte("updated description"))
}

func (u Users) Read(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		u.Log.ErrorHttp(http.StatusBadRequest, "failed to parse form", w)
		return
	}
	api := r.Form.Get("api_key")
	user, err := u.DB.RetriveUser(api)
	if err != nil {
		u.Log.ErrorHttp(http.StatusInternalServerError, "cannot retrive user session"+err.Error(), w)
		return
	}
	v, _ := json.Marshal(user)
	w.Write(v)
}

func (u Users) Regenerate(w http.ResponseWriter, r *http.Request) {
}

func (u Users) Error(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("error on user"))
}
