package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Student struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type Response struct {
	Object string    `json:"object"`
	Data   []Student `json:"data"`
}

type Server struct {
	db *sql.DB
}

func main() {
	db, err := connectDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = retry(5, 3*time.Second, db.Ping)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB!")

	log.Println("Seeding the database...")
	err = seedDb(db)
	if err != nil {
		log.Fatal(err)
	}

	server := NewServer(db)
	server.Run()
}

func retry(attempts int, timeout time.Duration, fn func() error) error {
	for i := 0; i <= attempts; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		log.Printf("Failed to connect to DB: %s\n", err)
		log.Printf("Retrying... attempt=%d, max_attempts=%d", i+1, attempts)
		time.Sleep(timeout)
	}
	return errors.New("unable to connect to DB")
}

func connectDb() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	portString := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, err
	}

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	log.Printf("Connecting: %s\n", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, err
}

func seedDb(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS students")
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name VARCHAR(160) NOT NULL,
		age INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	students := []Student{
		Student{
			Name: "Alyx",
			Age:  10,
		},
		Student{
			Name: "Bon",
			Age:  11,
		},
		Student{
			Name: "Ciel",
			Age:  12,
		},
	}
	var wg sync.WaitGroup
	for _, student := range students {
		wg.Add(1)
		go func(student Student) {
			_, err = db.Exec("INSERT INTO students(name, age) VALUES ($1, $2)", student.Name, student.Age)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(student)
	}
	wg.Wait()

	return nil
}

func NewServer(db *sql.DB) Server {
	return Server{db}
}

func (s *Server) Run() {
	http.HandleFunc("/", makeHandler(s.db))
	log.Println("Running on port 9000...")
	http.ListenAndServe(":9000", nil)
}

func makeHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, age, created_at FROM students")
		defer func() {
			err := rows.Close()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		students := []Student{}
		for rows.Next() {
			var student Student
			err = rows.Scan(&student.Id, &student.Name, &student.Age, &student.CreatedAt)
			if err != nil {
				log.Println(err)
			} else {
				students = append(students, student)
			}
		}

		payload := Response{
			Object: "students",
			Data:   students,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
