package main

import (
	"net/http"
)

type V1ErrorResponse struct {
	Error string `json:"error"`
}

func respondToError(sponse http.ResponseWriter, quest *http.Request) {
	respondWithJson(sponse, 500, V1ErrorResponse{
		Error: "Internal Server Error",
	})
}
