package logutil

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func SetLogger(logPath, logFileName string) {
	configLocalFilesystemLogger(logPath, logFileName, time.Hour*24*30, time.Hour*24)
}

func ErrorLog(args ...interface{}) {
	log.Error(args...)
}

func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}
