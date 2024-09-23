package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshparkerj/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

var config ApiConfig

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")
	fmt.Println("this is the main. Welcome!")
	fmt.Println("now I can edit the file yay!")
	fmt.Println(port)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)

		log.Fatal("database didn't open")
	}

	config = ApiConfig{
		DB: database.New(db),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /smiley", func(sponse http.ResponseWriter, quest *http.Request) {
		// sponse.Write([]byte("say cheese!"))
		err := respondWithJson(sponse, 201, GlobalTestPayload)
		if err != nil {
			fmt.Println("got an error")
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/smiley-error", func(sponse http.ResponseWriter, quest *http.Request) {
		err := respondWithError(sponse, 400, "this was an error")
		if err != nil {
			fmt.Println("got an error")
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/v1/healthz", readiness)
	mux.HandleFunc("/v1/err", respondToError)
	mux.HandleFunc("POST /v1/users", createUser)
	mux.HandleFunc("GET /v1/users", func(sponse http.ResponseWriter, quest *http.Request) {
		err := respondWithError(sponse, 500, "not implemented")
		if err != nil {
			fmt.Println("got an error")
			fmt.Println(err)
		}
	})

	server := http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%v", port),
	}

	server.ListenAndServe()
}
