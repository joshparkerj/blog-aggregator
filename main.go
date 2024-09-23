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
		respondWithJson(sponse, 200, GlobalTestPayload)
	})

	server := http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%v", port),
	}

	server.ListenAndServe()
}
