package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var bind = "0.0.0.0:8000"
var key = os.Getenv("key")

// InfoHandler - func.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare response struct.
	responseData := struct {
		Timestamp string `json:"timestamp"`
		Hostname  string `json:"hostname"`
	}{}

	// Get Timestamp string.
	responseData.Timestamp = fmt.Sprint(
		time.Now().Unix(),
	)
	// Get Hostname.
	responseData.Hostname, _ = os.Hostname()

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(responseData)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/info", InfoHandler).Methods("GET", "HEAD")

	fmt.Printf("Serving on %v...\n", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}
