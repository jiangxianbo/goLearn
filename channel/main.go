package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要指定通道中的元素
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)     // nil
	b = make(chan int) // 通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 // 卡住了
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)         // nil
	b = make(chan int, 10) // 通道初始化
	b <- 10                // 卡住了
	fmt.Println("10发送到通道b中了...")
	b <- 20
	fmt.Println("20发送到通道b中了...") // 指定
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)

}

func main() {
	//noBufChannel()
	bufChannel()
}
