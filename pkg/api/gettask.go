package api

import (
	"FINAL-PROJECT/pkg/db"
	"log"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработчик getTaskHandler")

	id := r.URL.Query().Get("id")
	log.Printf("Получен id: %s", id)
	task, err := db.GetTask(id)
	if err != nil {
		log.Printf("Ошибка при получении id: %v", err)
		returnError(w, err.Error())
		return
	}
	if err = writeJson(w, task); err != nil {
		returnError(w, "Task can not Marshal")
		return
	}
}
