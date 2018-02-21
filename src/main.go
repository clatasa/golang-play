package main

import (
	"log"
	"net/http"
	"os"
)

var (
	Info *log.Logger
)

func main() {

	defer LogFileSetup("server.log").Close()

	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}

func LogFileSetup(logFileName string) *os.File {
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	Info = log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return f
}
