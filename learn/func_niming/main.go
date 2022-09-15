package main

import "fmt"

//var f1 = func(x, y int) {
//	fmt.Println(x + y)
//}

func main() {
	// 匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果只是调用一次，立即执行函数
	func(x, y int) {
		fmt.Println("Hello World")
	}(100, 200)
}
