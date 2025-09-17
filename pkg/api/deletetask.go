package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"FINAL-PROJECT/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := strconv.Atoi(id)
	if id == "" || err != nil {
		returnError(w, fmt.Sprintf("ID is empty or %v", err), 400)
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		returnError(w, "не получилось удалить задачу", 503)
	}

	type EmptyStruct struct{}
	var emptyJson EmptyStruct
	answerEmptyJson, err := json.Marshal(emptyJson)
	if err != nil {
		returnError(w, err.Error(), 503)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(answerEmptyJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
