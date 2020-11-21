package logs

import (
	"log"
	"os"
)

// InfoLog logs the info gotten within the application
var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

// ErrorLog logs the error gotten within the application
var ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)
