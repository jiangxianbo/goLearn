package main

import "fmt"

func main() {
	s1 := make([]int, 1, 10)
	fmt.Printf("v:%v len(t):%v cap(t):%v\n", s1, len(s1), cap(s1))

	a11 := [...]int{1,3,5}
	x1 := a11[:]
	fmt.Printf("v:%v len(t):%v cap(t):%v\n", x1, len(x1), cap(x1))
	x1 = append(x1[:1], x1[2:]...)
	fmt.Println(a11)
}
