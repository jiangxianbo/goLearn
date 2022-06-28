package main

import "fmt"

func main() {
	n := 19
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

}
