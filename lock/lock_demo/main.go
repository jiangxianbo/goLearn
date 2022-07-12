package main

// 读写锁

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x = 0
var lock sync.Mutex

func add() {

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()

}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
