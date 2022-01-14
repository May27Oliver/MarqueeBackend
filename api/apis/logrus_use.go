package apis

import (
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logPath = "./log"
	logFile = "marquee.log"
)

var LogInstance = logrus.New()

func init() {
	logFileName := path.Join(logPath, logFile)
	rolling(logFileName)
	//log輸出格式
	LogInstance.SetFormatter(&logrus.JSONFormatter{})
	// log記錄級別
	LogInstance.SetLevel(logrus.DebugLevel)
}

func rolling(logFile string) {
	// 設置輸出
	LogInstance.SetOutput(&lumberjack.Logger{
		Filename:   logFile, //日誌文件位置
		MaxSize:    50,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	})
}
