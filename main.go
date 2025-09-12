package main

import (
	"FINAL-PROJECT/pkg/db"
	"FINAL-PROJECT/pkg/server"
	"log"
	"os"
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

	err = db.Init("scheduler.db")
	if err != nil {
		log.Printf("Mistake in db. ")
	}

	server.Run()

}
