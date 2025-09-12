package api

import (
	"FINAL-PROJECT/pkg/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработчик deleteTaskHandler")

	id := r.URL.Query().Get("id")
	log.Printf("Получен DELETE запрос на удаление, id: %s", id)

	_, err := strconv.Atoi(id)
	if id == "" || err != nil {
		log.Println("при проверке ID выявлена ошибка: пустой или неформат")
		returnError(w, fmt.Sprintf("ID is empty or %v", err))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		log.Printf("Ошибка при получении id: %v", err)
		returnError(w, err.Error())
		return
	}
	log.Printf("Через GetTask получена задача: %+v", task)

	err = db.DeleteTask(task.ID)
	if err != nil {
		returnError(w, "не получилось удалить задачу")
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
