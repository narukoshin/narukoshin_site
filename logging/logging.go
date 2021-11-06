package logging

import (
	"runtime"
	"time"
	"log"
	"os"
)

var logFolder string = "errors/"

func formatFilename() string {
	// getting the current date
	date := time.Now().Format("2006-01-02")
	// building file name and returning
	return date + "_errors.log"
}

func Init(){
	// building full path of the error log file
	filePath := logFolder + formatFilename()
	// Opening the error log file or creating if does not exist
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// setting the output in the error log file
	log.SetOutput(file)
}

func Save(message interface{}){
	// getting the file path where the error is
	_, filePath, _, _:= runtime.Caller(1)
	// printing error message to the file
	log.Print(message," in ", filePath)
}

func Savef(format string, v ...interface{}) {
	log.Printf(format, v...)
}