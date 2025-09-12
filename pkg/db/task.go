package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {

	var id int64
	log.Printf("Начинаем добавлять данные в базу данных/%s, /%s, /%s, /%s", task.Date, task.Title, task.Comment, task.Repeat)
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Printf("При открытии базы данных получена ошибка: %v", err)
		return 0, err
	}
	defer db.Close()

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`
	res, err := db.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))
	if err != nil {
		log.Printf("ошибка db.Exec - %v", err)
	}
	if err == nil {
		id, err = res.LastInsertId()
		log.Printf("ошибка res.LastInsertId - %v", err)
	}
	log.Printf("ID равен %d", id)
	return id, err
}

func Tasks(limit int) ([]*Task, error) {
	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Printf("При открытии базы данных получена ошибка: %v", err)
		return nil, err
	}
	defer db.Close()

	resSlice := make([]*Task, 0)

	rows, err := db.Query("SELECT * FROM scheduler ORDER BY date LIMIT :limit", sql.Named("limit", limit))
	if err != nil {
		log.Println("Ошибка при SELECT запросе в БД")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Println("Ошибка сканирования запроса в БД")
			return nil, err
		}
		resSlice = append(resSlice, &task)
	}
	log.Printf("После SELECT запроса получили слайс структур: %+v", resSlice)

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return resSlice, nil
}

func GetTask(id string) (*Task, error) {

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Printf("При открытии базы данных получена ошибка: %v", err)
		return nil, err
	}
	defer db.Close()

	task := Task{}

	row := db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id))

	err = row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		log.Println("Ошибка сканирования запроса в БД")
		return nil, err
	}
	log.Printf("После SELECT запроса для редактирования получили структуру: %+v", task)

	return &task, nil
}

func UpdateTask(task *Task) error {
	log.Println("функция UpdateTask")
	log.Printf("структура task: %+v", task)

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Printf("При открытии базы данных получена ошибка: %v", err)
		return err
	}
	defer db.Close()

	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`
	res, err := db.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat), sql.Named("id", task.ID))
	if err != nil {
		log.Println("при отправке запроса Exec на изменение произошла ошибка")
		return err
	}
	log.Println("отправлен Exec запрос на изменение")
	// метод RowsAffected() возвращает количество записей к которым
	// был применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("проверка 	count, err := res.RowsAffected() прошла успешно")
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	log.Println("проверка 	count == 0 прошла успешно")
	log.Println("функция UpdateTask вернула nil")
	return nil
}

func UpdateDate(next string, id string) error {
	log.Println("функция UpdateDate")

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Printf("При открытии базы данных получена ошибка: %v", err)
		return err
	}
	defer db.Close()

	query := `UPDATE scheduler SET date = :date  WHERE id = :id`
	res, err := db.Exec(query, sql.Named("date", next), sql.Named("id", id))
	if err != nil {
		log.Println("при отправке запроса Exec на изменение произошла ошибка")
		return err
	}
	log.Println("отправлен Exec запрос на изменение даты")
	// метод RowsAffected() возвращает количество записей к которым
	// был применена SQL команда
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("проверка 	count, err := res.RowsAffected() прошла успешно")
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	log.Println("проверка 	count == 0 прошла успешно")
	log.Println("функция UpdateTask вернула nil")
	return nil
}

func DeleteTask(id string) error {
	log.Println("функция DeleteTask")

	db, err := sql.Open("sqlite", "scheduler.db")
	if err != nil {
		log.Printf("При открытии базы данных получена ошибка: %v", err)
		return err
	}
	defer db.Close()

	query := `DELETE FROM scheduler  WHERE id = :id`
	_, err = db.Exec(query, sql.Named("id", id))
	if err != nil {
		log.Println("при отправке запроса DELETE произошла ошибка")
		return err
	}
	log.Println("отправлен запрос на удаление - успешно")
	return nil
}
