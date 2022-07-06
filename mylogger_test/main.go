package main

import (
	"goLearn/mylogger"
	"time"
)

// 测试自己写的日志库
func main() {
	log := mylogger.NewLog("Error")
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条console日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}

}
