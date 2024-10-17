package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	"go_final_project/database"
	"go_final_project/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}

func fileServer(path string) http.Handler {
	fs := http.FileServer(http.Dir(path))
	return fs
}

func main() {

	db, err := database.NewDatabase()
	if err != nil {
		log.Printf("failed to set up database: %v", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()
	log.Println("Database setup successfully")

	portStr := os.Getenv("TODO_PORT")
	var port int

	if portStr == "" {
		port = 7540
	} else {
		port, err = strconv.Atoi(portStr)
		if err != nil {
			log.Printf("Invalid port number: %v", err)
		}
	}

	webDir := "./web"

	r := chi.NewRouter()

	r.Handle("/*", fileServer(webDir))
	log.Printf("Loaded frontend files from %s\n", webDir)

	r.Get("/api/nextdate", handlers.GetNextDate)
	r.Post("/api/task", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTaskHandler(w, r, db)
	})
	r.Get("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r, db)
	})
	r.Get("/api/task", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTaskByIDHandler(w, r, db)
	})
	r.Put("/api/task", func(w http.ResponseWriter, r *http.Request) {
		handlers.EditTaskHandler(w, r, db)
	})
	r.Post("/api/task/done", func(w http.ResponseWriter, r *http.Request) {
		handlers.DoneTaskHandler(w, r, db)
	})
	r.Delete("/api/task", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(w, r, db)
	})

	address := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on http://localhost%s\n", address)
	if err := http.ListenAndServe(address, r); err != nil {
		log.Printf("Server failed to start: %v\n", err)
	}
}
