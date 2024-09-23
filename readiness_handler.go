package main

import (
	"net/http"
)

type ReadinessResponse struct {
	Status string `json:"status"`
}

func readiness(sponse http.ResponseWriter, quest *http.Request) {
	respondWithJson(sponse, 200, ReadinessResponse{
		Status: "ok",
	})
}
