package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req HelloRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		resp := HelloResponse{
			Message: "Hello " + req.Name,
			Time:    time.Now().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	if r.Method == http.MethodGet {
		resp := HelloResponse{
			Message: "Hello World",
			Time:    time.Now().Format(time.RFC3339),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	http.Error(w, "Unrecognised Method", http.StatusMethodNotAllowed)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Listening on :8080")
	log.Fatal(server.ListenAndServe())
}
