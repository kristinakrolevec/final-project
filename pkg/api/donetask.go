package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"FINAL-PROJECT/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	_, err := strconv.Atoi(id)
	if id == "" || err != nil {
		returnError(w, fmt.Sprintf("ID is empty or %v", err), 404)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		returnError(w, err.Error(), 404)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(task.ID)
		if err != nil {
			returnError(w, "не получилось удалить задачу", 503)
		}

	} else {
		task.Date, err = NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			returnError(w, err.Error(), 404)
			return
		}
		err = db.UpdateDate(task.Date, task.ID)
		if err != nil {
			returnError(w, err.Error(), 503)
			return
		}
	}

	type EmptyStruct struct{}
	var emptyJson EmptyStruct
	answerEmptyJson, err := json.Marshal(emptyJson)
	if err != nil {
		returnError(w, err.Error(), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(answerEmptyJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
