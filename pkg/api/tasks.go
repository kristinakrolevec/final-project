package api

import (
	"net/http"

	"FINAL-PROJECT/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

const limit = 50

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	tasks, err := db.Tasks(limit)
	if err != nil {
		returnError(w, err.Error(), 404)
		return
	}
	if err = writeJson(w, TasksResp{
		Tasks: tasks,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
