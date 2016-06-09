//	PhalGo-Logs
//	日志记录功能
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/Sirupsen/logrus"  基于logrus的log类

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

// 初始化Log日志记录
func init() {
	var err error
	LogS = logrus.New()
	LogS.Formatter = new(logrus.JSONFormatter)
	//log.Formatter = new(logrus.TextFormatter) // default
	LogS.Level = logrus.DebugLevel

	if !IsDirExists(GetPath() + "/Runtime") {
		if mkdirerr := MkdirFile(GetPath() + "/Runtime"); mkdirerr != nil {
			fmt.Println(mkdirerr)
		}
	}

	logfile, err = os.OpenFile(GetPath() + "/Runtime/" + time.Now().Format("2006-01-02") + ".log", os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		logfile, err = os.Create(GetPath() + "/Runtime/" + time.Now().Format("2006-01-02") + ".log")
		if err != nil {
			fmt.Println(err)
		}
	}
	LogS.Out = logfile
	day = time.Now().Format("02")
}

// 检测是否跨天了,把记录记录到新的文件目录中
func updateLogFile() {
	var err error
	day2 := time.Now().Format("02")
	if day2 != day {
		logfile.Close()
		logfile, err = os.Create(GetPath() + "/Runtime/" + time.Now().Format("2006-01-02") + ".log")
		if err != nil {
			fmt.Println(err)
		}
		LogS.Out = logfile
	}
}

// 记录Debug信息
func LogDebug(str interface{}, data logrus.Fields) {
	updateLogFile()
	LogS.WithFields(data).Debug(str)
}

// 记录Info信息
func LogInfo(str interface{}, data logrus.Fields) {
	updateLogFile()
	LogS.WithFields(data).Info(str)
}

// 记录Error信息
func LogError(str interface{}, data logrus.Fields) {
	updateLogFile()
	LogS.WithFields(data).Error(str)
}
