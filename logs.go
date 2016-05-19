package phalgo

//	PhalGo-Logs
//	日志记录功能
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/Sirupsen/logrus"  基于logrus的log类

import (
	"github.com/Sirupsen/logrus"
	"os"
	"time"
	"fmt"
)

var LogS *logrus.Logger

var day string
var logfile *os.File

func init() {
	var err error
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
	var err error
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

func LogDebug(str interface{}, data logrus.Fields) {
	updateLog()
	LogS.WithFields(data).Debug(str)
}

func LogInfo(str interface{}, data logrus.Fields) {
	updateLog()
	LogS.WithFields(data).Info(str)
}

func LogError(str interface{}, data logrus.Fields) {
	updateLog()
	LogS.WithFields(data).Error(str)
}
