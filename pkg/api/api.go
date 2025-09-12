package api

import (
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.HandleFunc("/api/nextdate", nextDayHandler)
	mux.HandleFunc("POST /api/task", addTaskHandler)
	mux.HandleFunc("GET /api/task", getTaskHandler)
	mux.HandleFunc("PUT /api/task", updateTaskHandler)
	mux.HandleFunc("GET /api/tasks", tasksHandler)
	mux.HandleFunc("POST /api/task/done", doneTaskHandler)
	mux.HandleFunc("DELETE /api/task", deleteTaskHandler)
}
