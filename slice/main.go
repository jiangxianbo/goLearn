package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	//
}
