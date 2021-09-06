package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

var id string

func getId() string {
	return id
}

func main() {
	// Init random ID
	rand.Seed(time.Now().UnixNano())
	id = strconv.Itoa(rand.Intn(1000))

	http.HandleFunc("/", rootHandler)
	fmt.Printf("[%s] Running on port 9000...\n", getId())
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	id := getId()
	log.Printf("[%s] %s %s\n", id, r.Method, r.RequestURI)
	payload := Response{
		Message:   fmt.Sprintf("[%s] Hello!", id),
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
