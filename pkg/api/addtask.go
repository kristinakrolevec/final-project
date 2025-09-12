package api

import (
	"FINAL-PROJECT/pkg/db"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	var next string

	if task.Date == "" || task.Date == now.Format("20060102") {
		task.Date = now.Format(DateFormat)
		log.Printf("Date в Post запросе пустой или сегодня. Берем сегодня: %v/", task.Date)
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
			log.Printf("task.Date меньше настоящей и Repeat = пусто, поэтому сегодня = %v", task.Date)
		} else {
			task.Date = next
			log.Printf("task.Date меньше настоящей и Repeat не пусто, поэтому сипользуем функцию NextDate и получаем: task.Date = %v", task.Date)
		}
	}
	return nil
}

func writeJson(w http.ResponseWriter, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
	log.Println("Записан заголовок и ответ после сериализации")
	return nil
}

func returnError(w http.ResponseWriter, errorMessage string) {
	errorResponse := map[string]string{"error": errorMessage}
	jsonResponse, _ := json.Marshal(errorResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработчик addTaskHandler")

	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		log.Printf("Ошибка при чтении запроса Post: %v/", err)
		returnError(w, err.Error())
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {

		log.Printf("Ошибка при десериализации в запросе Post: %v/", err)
		returnError(w, err.Error())
		return
	}
	log.Printf("Получен POST запрос в виде: %+v", task)

	if task.Title == "" {
		log.Println("Title в Post запросе пустой")
		returnError(w, "Title is empty")
		return
	}

	if err = checkDate(&task); err != nil {
		log.Println("Date format error")
		returnError(w, "Date format is error")
		return
	}
	id, err := db.AddTask(&task)
	if err != nil {
		log.Println("Error creating ID")
		returnError(w, "Error creating ID")
		return
	}

	log.Printf("присвоенный ID равен %d", id)
	task.ID = strconv.Itoa(int(id))

	idResponse := map[string]string{"id": task.ID}

	if err = writeJson(w, idResponse); err != nil {
		returnError(w, "ID can not Marshal")
		return
	}

}
