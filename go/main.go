package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ssse-exercise-sieve/pkg/sieve"
)

func main() {
	http.HandleFunc("/prime", primeNumHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

type payload struct {
	NthPrime int64 `json:"nth_prime"`
}

func primeNumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if r.URL.Path != "/prime" {
		http.Error(w, "invalid request uri", http.StatusNotFound)
	}

	// Decode JSON request body
	var data payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)

	s := sieve.NewSieve()
	output := s.NthPrime(data.NthPrime)

	err = json.NewEncoder(w).Encode(map[string]int64{"Prime Number": output})
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
}
