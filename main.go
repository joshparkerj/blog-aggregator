package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshparkerj/blog-aggregator/internal/config"
	"github.com/joshparkerj/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

var state State
var commands Commands

func main() {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	configuration, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	state.Configuration = &configuration

	commands.Commands = make(map[string]func(*State, Command) error)

	// register a handler function for the login command
	commands.Register("login", Login)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments!")
	}

	commandName := args[1]
	args = args[2:]
	err = commands.Run(&state, Command{
		Name: commandName,
		Args: args,
	})

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", dbURL)

	state.DB = database.New(db)

	if err != nil {
		fmt.Println(err)

		log.Fatal("database didn't open")
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
}
