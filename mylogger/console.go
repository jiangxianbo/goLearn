package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// Newlog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARING) {
		log(WARING, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
