package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(x, y int, f func(int, int)) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

//func f3(f func(int, int), m, n int) func() {
//	tmp := func() {
//		f(m, n)
//	}
//	return tmp
//}
func main() {
	//f1(f3(f2, 1, 2))
	f1(f3(1, 2, f2))
}
