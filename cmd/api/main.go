package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Second,
	}

	log.Println("Starting server :8080")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}

}

func hello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, Go!")
}
