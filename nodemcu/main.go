package main

import (
	"fmt"
	"log"
	"net/http"
)

func uploadDataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	dht11 := r.URL.Query().Get("dht11")
	switch1 := r.URL.Query().Get("switch1")
	switch2 := r.URL.Query().Get("switch2")

	// Print or process the data
	fmt.Printf("Received Data - DHT11: %s, Switch1: %s, Switch2: %s\n", dht11, switch1, switch2)

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))
}

func main() {
	http.HandleFunc("/upload", uploadDataHandler)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
