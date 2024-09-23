package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println("this is the main. Welcome!")
	fmt.Println("now I can edit the file yay!")
	fmt.Println(port)

	// now let's do some real coding!

	mux := http.NewServeMux()

	mux.HandleFunc("/smiley", func(sponse http.ResponseWriter, quest *http.Request) {
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

	server := http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%v", port),
	}

	server.ListenAndServe()
}
