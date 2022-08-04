package main

import (
	"fmt"
	"time"
)

// goroutine
func hello(i int) {
	fmt.Println("hello", i)
}

func main() {
	for i := 0; i < 10; i++ {
		//go hello(i) // 开始一个单独的goroutine去执行hello函数（任务）
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main3")
	// main函数结束了 由main函数启动的goroutine也都结束了
	time.Sleep(time.Second * 1)
}
