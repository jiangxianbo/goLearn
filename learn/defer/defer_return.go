package main

import "fmt"

// Go语言中的函数return不是原子操作，在底层是分为两步
// 第一：返回值赋值
// 第二：真正的RET返回
// 函数中如果存在defer,那么defer执行的时机是第一步和第二步之间

func f1() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}
