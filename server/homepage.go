package server

import (
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Janki"))
}
