package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"FINAL-PROJECT/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		returnError(w, err.Error(), 400)
		return
	}
	_, err = strconv.Atoi(task.ID)
	if task.ID == "" || err != nil {
		returnError(w, fmt.Sprintf("ID is empty or %v", err), 400)
		return
	}
	_, err = db.GetTask(task.ID)
	if err != nil {
		returnError(w, err.Error(), 404)
		return
	}

	if task.Title == "" {
		returnError(w, "Title is empty", 400)
		return
	}

	if err = checkDate(&task); err != nil {
		returnError(w, "Date format is error", 400)
		return
	}
	err = db.UpdateTask(&task)
	if err != nil {
		returnError(w, "Error update task", 503)
		return
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
	if _, err = w.Write(answerEmptyJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
