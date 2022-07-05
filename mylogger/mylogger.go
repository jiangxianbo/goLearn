package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

const (
	// 定义日志级别
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "waring":
		return WARING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARING:
		return "WARING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller failed, err:%v\n")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
