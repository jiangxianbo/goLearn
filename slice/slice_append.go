package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(a1)
	s1 := a1[:]
	s1 = append(s1[:1], s1[4:]...)
	fmt.Println(s1)
	fmt.Println(a1)
}
