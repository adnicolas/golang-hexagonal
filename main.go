package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddr = ":8081"

func main() {
	fmt.Println("Server running on", httpAddr)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

// Underscore convention to indicate that I'm not going to use that parameter
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("everything is ok!"))
}
