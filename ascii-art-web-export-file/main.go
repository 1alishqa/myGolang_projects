package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public/", files))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/process", processHandler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting Server On Port 8080 ...")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		os.Exit(1)
	}
	fmt.Println()
}
