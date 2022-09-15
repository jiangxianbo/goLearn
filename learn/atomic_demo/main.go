package main

import (
	"fmt"
	"sync/atomic"
)

var x int64 = 100

func main() {
	a := atomic.LoadInt64(&x)
	fmt.Println(a)
}
