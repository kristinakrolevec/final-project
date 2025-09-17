package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"FINAL-PROJECT/pkg/db"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	var next string

	if task.Date == "" || task.Date == now.Format(DateFormat) {
		task.Date = now.Format(DateFormat)
		return nil
	}

	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		log.Println("Invalid date format")
		return err
	}
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			log.Printf("получили ошибку после NextDate: %v", err)
			return err
		}
	}

	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(DateFormat)
		} else {
			task.Date = next
		}
	}
	return nil
}

func writeJson(w http.ResponseWriter, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %v", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func returnError(w http.ResponseWriter, errorMessage string, status int) {
	errorResponse := map[string]string{"error": errorMessage}
	jsonResponse, _ := json.Marshal(errorResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(jsonResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s", errorMessage)
	}
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

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

	if task.Title == "" {
		returnError(w, "Title is empty", 400)
		return
	}

	if err = checkDate(&task); err != nil {
		returnError(w, "Date format is error", 400)
		return
	}
	id, err := db.AddTask(&task)
	if err != nil {
		returnError(w, "Error creating ID", 503)
		return
	}

	task.ID = strconv.Itoa(int(id))

	idResponse := map[string]string{"id": task.ID}

	if err = writeJson(w, idResponse); err != nil {
		returnError(w, "ID can not Marshal", 503)
		return
	}

}
