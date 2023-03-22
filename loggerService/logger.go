package loggerService

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {
	//set location to the log file
	file, err := os.OpenFile("./blogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
