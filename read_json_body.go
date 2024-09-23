package main

import (
	"encoding/json"
	"net/http"
)

func readJsonBody[T any](quest *http.Request, emptyStruct T) (payload T, err error) {
	decoder := json.NewDecoder(quest.Body)
	err = decoder.Decode(&emptyStruct)
	payload = emptyStruct
	return
}
