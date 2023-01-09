package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/pab-framework/tool"
	"io"
	"os"
	"strings"
	"time"
)

type CatLog struct {
	logOption
}

type logOption struct {
	timeFormat      string
	level           int
	callerFormat    string
	goroutineFormat string
	logFile         *lumberjack.Logger
	outLog          func(w io.Writer, msg string)
}

type Option func(*logOption)

func NewCatLog(opts ...Option) Log {
	option := newDefaultLogOption(opts...)
	cl := CatLog{
		logOption: option,
	}
	return &cl
}

func WithTimeFormat(format string) Option {
	return func(option *logOption) {
		if strings.EqualFold(format, "") {
			return
		}
		option.timeFormat = format
	}
}

func WithLevel(level int) Option {
	return func(option *logOption) {
		if level <= 0 {
			return
		}
		option.level = level
	}
}

func WithCaller(caller string) Option {
	return func(option *logOption) {
		if strings.EqualFold(caller, "") {
			return
		}
		option.callerFormat = caller
	}
}

func WithGoroutine(goroutine string) Option {
	return func(option *logOption) {
		if strings.EqualFold(goroutine, "") {
			return
		}
		option.goroutineFormat = goroutine
	}
}

func WithLogFile(logPath string) Option {
	return WithLogFileParam(logPath, 500, 20, 7, false)
}

func WithLogFileParam(logPath string, maxSize, maxAge, maxBackups int, enableCompress bool) Option {
	return func(option *logOption) {
		if strings.EqualFold(logPath, "") {
			return
		}
		option.logFile = &lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    maxSize,    // 日志大小
			MaxAge:     maxAge,     // 保留数量
			MaxBackups: maxBackups, // 保留天数
			Compress:   enableCompress,
		}
		option.outLog = outFile
	}
}

func newDefaultLogOption(opts ...Option) logOption {
	defaultOption := logOption{
		timeFormat:      "2006-01-02T15:04:05.000",
		level:           InfoLevel,
		callerFormat:    "[%s]",
		goroutineFormat: "[_it=%d]",
		outLog:          outConsole,
	}
	for _, apply := range opts {
		apply(&defaultOption)
	}
	return defaultOption
}

func outConsole(w io.Writer, msg string) {
	fmt.Println(msg)
}

func outFile(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func (cl *CatLog) printLog(level int, caller, msg string) {
	var l string
	switch level {
	case DebugLevel:
		l = "[DEBUG]"
	case TraceLevel:
		l = "[TRACE]"
	case InfoLevel:
		l = "[INFO]"
	case WarnLevel:
		l = "[WARN]"
	case ErrorLevel:
		l = "[ERROR]"
	case FatalLevel:
		l = "[FATAL]"
	default:
		l = "[INFO]"
	}
	if level >= cl.level {
		cl.outLog(cl.logFile, fmt.Sprintf("%s %s %s %d -- %s\n", time.Now().Format(cl.timeFormat), l, fmt.Sprintf(cl.callerFormat, caller), tool.GetGoroutineID(), msg))
	}
}

func (cl *CatLog) IPrint(level int, caller, msg string) {
	cl.printLog(level, caller, msg)
}

func (cl *CatLog) IPrintF(level int, caller, format string, args ...interface{}) {
	cl.printLog(level, caller, fmt.Sprintf(format, args...))
}

func (cl *CatLog) IFatal(level int, caller string, err error) {
	cl.printLog(level, caller, err.Error())
	os.Exit(1)
}
