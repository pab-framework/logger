package logger

import (
	"errors"
	"fmt"
)

var log Log

func InitLog(logPath string) {
	log = NewCatLog(WithLogFile(logPath))
}

func Debug(msg string) {
	log.IPrint(DebugLevel, catCaller(), msg)
}

func DebugF(format string, args ...interface{}) {
	log.IPrintF(DebugLevel, catCaller(), format, args...)
}

func Trace(msg string) {
	log.IPrint(TraceLevel, catCaller(), msg)
}

func TraceF(format string, args ...interface{}) {
	log.IPrintF(TraceLevel, catCaller(), format, args...)
}

func Info(msg string) {
	log.IPrint(InfoLevel, catCaller(), msg)
}

func InfoF(format string, args ...interface{}) {
	log.IPrintF(InfoLevel, catCaller(), format, args...)
}

func Warn(msg string) {
	log.IPrint(WarnLevel, catCaller(), msg)
}

func WarnF(format string, args ...interface{}) {
	log.IPrintF(WarnLevel, catCaller(), format, args...)
}

func Error(msg string) {
	log.IPrint(ErrorLevel, catCaller(), msg)
}

func ErrorF(format string, args ...interface{}) {
	log.IPrintF(ErrorLevel, catCaller(), format, args...)
}

func Fatal(err error) {
	log.IFatal(FatalLevel, catCaller(), err)
}

func FatalF(format string, args ...interface{}) {
	log.IFatal(FatalLevel, catCaller(), errors.New(fmt.Sprintf(format, args...)))
}

func catCaller() string {
	return caller(3)
}
