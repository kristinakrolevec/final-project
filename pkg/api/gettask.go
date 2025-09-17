package api

import (
	"net/http"

	"FINAL-PROJECT/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	task, err := db.GetTask(id)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}
	if err = writeJson(w, task); err != nil {
		returnError(w, "Task can not Marshal", 503)
		return
	}
}
