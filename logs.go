package phalgo

import (
	"github.com/Sirupsen/logrus"
	"os"
	"time"
	"fmt"
)

var LogS *logrus.Logger

var day string
var logfile *os.File
var err error


func init() {
	LogS = logrus.New()
	LogS.Formatter = new(logrus.JSONFormatter)
	//log.Formatter = new(logrus.TextFormatter) // default
	LogS.Level = logrus.DebugLevel
	logfile, err = os.OpenFile("Runtime/" + time.Now().Format("2006-01-02") + ".log", os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		logfile, err = os.Create("Runtime/" + time.Now().Format("2006-01-02") + ".log")
		if err != nil {
			fmt.Println(err)
		}
	}
	LogS.Out = logfile
	day = time.Now().Format("02")
}

func updateLog() {
	day2 := time.Now().Format("02")
	if day2 != day {
		logfile.Close()
		logfile, err = os.Create("Runtime/" + time.Now().Format("2006-01-02") + ".log")
		if err != nil {
			fmt.Println(err)
		}
		LogS.Out = logfile
	}
}

func Debug(str string, data logrus.Fields) {
	updateLog()
	LogS.WithFields(data).Debug(str)
}

func Error(str interface{}, data logrus.Fields) {
	updateLog()
	LogS.WithFields(data).Error(str)
}
