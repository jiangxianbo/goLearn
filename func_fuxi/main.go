package main

import "fmt"

func f1(a ...interface{}) {
	fmt.Printf("type:%T，value:%v\n", a, a)
}

func main() {
	//f1()
	//f1(1)
	//f1(1, 2, 3, 4, 5)
	//f1(false, true, [3]int{1, 23, 4}, map[string]int{"x": 1, "y": 2})
	var s = []interface{}{}
	s = []interface{}{1, 2, 3, 4, 5, 6}
	f1(s)
	f1(s...)
}
