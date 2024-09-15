package server

import (
	"janki/api"
	"net/http"
)

func Handler(api *api.Api) *http.ServeMux {
	h := http.NewServeMux()

	h.HandleFunc("/users/create", api.Users.Create)
	h.HandleFunc("/users/read", api.Users.Read)
	h.HandleFunc("/users/", api.Users.Error)

	h.HandleFunc("/knob/create", api.Knob.Create)

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Your error knows no bounds"))
	})

	return h
}
