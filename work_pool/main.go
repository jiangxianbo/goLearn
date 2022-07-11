package main

import (
	"fmt"
	"time"
)

var notifyCh = make(chan struct{}, 5) // 定义通知通道

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start jod:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end jod:%d\n", id, j)
		results <- j * 2
		notifyCh <- struct{}{}
	}

}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 5个任务
	go func() {
		for j := 0; j < 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// 开启三个goroutine
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	// 取通知
	go func() {
		for i := 0; i < 5; i++ {
			<-notifyCh
		}
		close(results)
	}()

	// 输出结果
	for x := range results {
		fmt.Println(x)
	}
}
