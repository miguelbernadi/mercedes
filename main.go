package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	// Default route
	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, req *http.Request) {
			_, err := io.WriteString(w, "Hello!")
			if err != nil {
				log.Println(err)
			}
		},
	)

	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
