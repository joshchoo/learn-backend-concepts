package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	http.HandleFunc("/", makeRootHandler(rdb))
	fmt.Printf("Running on port 9000...\n")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}

func makeRootHandler(rdb *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var count int
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&count); err != nil {
			sendError(w, err)
			return
		}

		var result int

		val, err := rdb.Get(r.Context(), getKey(count)).Result()
		if err == redis.Nil {
			result = fibonacci(count)
			expiration := 0 * time.Second
			if err = rdb.SetNX(r.Context(), getKey(count), result, expiration).Err(); err != nil {
				sendError(w, err)
				return
			}
		} else if err != nil {
			sendError(w, err)
			return
		} else {
			result, err = strconv.Atoi(val)
			if err != nil {
				sendError(w, err)
				return
			}
		}

		sendOk(w, result)
	}
}

func fibonacci(count int) int {
	if count <= 1 {
		return count
	}
	return fibonacci(count-1) + fibonacci(count-2)
}

func getKey(count int) string {
	return fmt.Sprintf("fib::%d", count)
}

func sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	if err2 := json.NewEncoder(w).Encode(err); err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func sendOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
