package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("你好！")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:

		}
	}
}

func worker2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("他好好！")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:

		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
