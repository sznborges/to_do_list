package main

import (
	"net/http"

	"github.com/sznborges/to_do_list/handlers"
)

func main() {
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.ListenAndServe(":8080", nil)
}

