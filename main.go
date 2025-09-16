package main

import (
	"FINAL-PROJECT/pkg/db"
	"FINAL-PROJECT/pkg/server"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Loggin to a file in GO!")

	dbPath := os.Getenv("TODO_DBFILE")
	if dbPath == "" {
		dbPath = "scheduler.db"
	}
	err = db.Init(dbPath)
	if err != nil {
		log.Printf("Mistake in db. ")
		return
	}
	defer db.Dbase.Close()

	server.Run()

}
