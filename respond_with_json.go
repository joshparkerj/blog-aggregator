package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload any) (err error) {
	fmt.Println(payload)
	marshalled, err := json.Marshal(payload)
	if err != nil {
		// if there was an error when trying to marshal the payload
		// then we will respond with a 500 code
		// before returning the error message to the caller
		// But should we handle this differently?
		code = 500
	}

	fmt.Println(string(marshalled))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(marshalled)
	return
}
