package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joshparkerj/blog-aggregator/internal/database"
)

type CreateUserReqRes struct {
	Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create user")
	emptyStruct := CreateUserReqRes{}
	payload, err := readJsonBody(r, emptyStruct)
	if err != nil {
		fmt.Println(err)
	}

	id := uuid.New()
	createUserParams := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      payload.Name,
	}
	state.DB.CreateUser(context.TODO(), createUserParams)
	respondWithJson(w, 201, payload)
}
