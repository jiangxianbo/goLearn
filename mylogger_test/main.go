package main

import (
	"goLearn/mylogger"
)

var log mylogger.Logger // 声明全局接口变量

// 测试自己写的日志库
func main() {
	log = mylogger.NewLConsoleLog("Error")                                   // 终端日志
	log = mylogger.NewFileLogger("Info", "./", "zhouling.log", 10*1024*1024) // 文件日志
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志, id: %d, name: %s", id, name)
		log.Fatal("这是一条Fatal日志")
		//time.Sleep(time.Second * 2)
	}

}
