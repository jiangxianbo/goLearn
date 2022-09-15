package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS

func A() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func B() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(16)
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}
