package api

import (
	"FINAL-PROJECT/pkg/db"
	"log"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Обработчик tasksHandler")

	tasks := make([]*db.Task, 0)
	log.Printf("Инициализирован слайс в структуре: %v", tasks)

	var err error
	tasks, err = db.Tasks(50)
	if err != nil {
		returnError(w, err.Error())
		return
	}
	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
