package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `CREATE TABLE scheduler (
id INTEGER PRIMARY KEY AUTOINCREMENT,
date CHAR(8) NOT NULL DEFAULT "",
title VARCHAR(256) NOT NULL DEFAULT "",
comment TEXT,
repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX scheduler_date ON scheduler (date);`

func Init(dbFile string) error {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}
	defer db.Close()
	return nil
}
