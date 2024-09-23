package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload any) (err error) {
	marshalled, err := json.Marshal(payload)
	if err != nil {
		code = 500
		w.Write([]byte("error"))
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(marshalled)
	return
}
