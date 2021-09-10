package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func getHostname() string {
	hn, err := os.Hostname()
	if err != nil {
		hn = "unknown"
	}
	return hn
}

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Printf("[%s] Running on port 9000...\n", getHostname())
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	hostname := getHostname()
	log.Printf("[%s] %s %s\n", hostname, r.Method, r.RequestURI)
	payload := Response{
		Message:   fmt.Sprintf("[%s] Hello!", hostname),
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
