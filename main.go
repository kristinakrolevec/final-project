package main

import (
	"log"
	"os"

	"FINAL-PROJECT/pkg/db"
	"FINAL-PROJECT/pkg/server"

	_ "modernc.org/sqlite"
)

func main() {
	log.SetOutput(os.Stdout)

	dbPath := os.Getenv("TODO_DBFILE")
	if dbPath == "" {
		dbPath = "scheduler.db"
	}
	err := db.Init(dbPath)
	if err != nil {
		log.Printf("Mistake in db. ")
		return
	}
	defer db.Dbase.Close()

	server.Run()

}
