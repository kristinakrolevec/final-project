package api

import (
	"FINAL-PROJECT/pkg/db"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработчик updateTaskHandler")

	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		log.Printf("Ошибка при чтении запроса Put: %v/", err)
		returnError(w, err.Error())
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {

		log.Printf("Ошибка при десериализации в запросе Put: %v/", err)
		returnError(w, err.Error())
		return
	}
	log.Printf("Получен PUT запрос в виде: %+v", task)

	_, err = strconv.Atoi(task.ID)
	if task.ID == "" || err != nil {
		log.Println("при проверке ID выявлена ошибка: пустой или неформат")
		returnError(w, fmt.Sprintf("ID is empty or %v", err))
		return
	}

	_, err = db.GetTask(task.ID)
	if err != nil {
		log.Printf("Ошибка при получении id: %v", err)
		returnError(w, err.Error())
		return
	}

	if task.Title == "" {
		log.Println("Title в Put запросе пустой")
		returnError(w, "Title is empty")
		return
	}

	if err = checkDate(&task); err != nil {
		log.Println("Date format error")
		returnError(w, "Date format is error")
		return
	}
	err = db.UpdateTask(&task)
	if err != nil {
		log.Println("Error update task")
		returnError(w, "Error update task")
		return
	}

	type EmptyStruct struct{}
	var emptyJson EmptyStruct
	answerEmptyJson, err := json.Marshal(emptyJson)
	if err != nil {
		log.Println("Ошибка сериализации пустой структуры")
		returnError(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(answerEmptyJson)
	w.WriteHeader(http.StatusOK)
	log.Println("Записан заголовок и ответ {} после сериализации")

}
