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

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)`
	res, err := Dbase.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))
	if err != nil {
		log.Printf("ошибка db.Exec - %v", err)
	}
	if err == nil {
		id, err = res.LastInsertId()
	}
	return id, err
}

func Tasks(limit int) ([]*Task, error) {

	resSlice := make([]*Task, 0)

	rows, err := Dbase.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit", sql.Named("limit", limit))
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
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return resSlice, nil
}

func GetTask(id string) (*Task, error) {

	task := Task{}

	row := Dbase.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id", sql.Named("id", id))

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		log.Println("Ошибка сканирования запроса в БД")
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *Task) error {

	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`
	res, err := Dbase.Exec(query, sql.Named("date", task.Date), sql.Named("title", task.Title), sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat), sql.Named("id", task.ID))
	if err != nil {
		log.Println("при отправке запроса Exec на изменение произошла ошибка")
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func UpdateDate(next string, id string) error {

	query := `UPDATE scheduler SET date = :date  WHERE id = :id`
	res, err := Dbase.Exec(query, sql.Named("date", next), sql.Named("id", id))
	if err != nil {
		log.Println("при отправке запроса Exec на изменение произошла ошибка")
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func DeleteTask(id string) error {

	query := `DELETE FROM scheduler  WHERE id = :id`
	_, err := Dbase.Exec(query, sql.Named("id", id))
	if err != nil {
		log.Println("при отправке запроса DELETE произошла ошибка")
		return err
	}
	return nil
}
