package logger

import (
	"fmt"
	"runtime"
)

const (
	DebugLevel = iota
	TraceLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type Log interface {
	IPrint(level int, caller, msg string)
	IPrintF(level int, caller, format string, args ...interface{})
	IFatal(level int, caller string, err error)
}

func caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	return fmt.Sprintf("[%s:%d]", file, line)
}
