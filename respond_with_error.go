package main

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	// NOTE: Do not forget to export all the fields that must be included in a json
	Msg string `json:"msg"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) (err error) {
	sponse := ErrorResponse{
		msg,
	}

	fmt.Println(sponse)

	err = respondWithJson(w, code, sponse)
	return
}
