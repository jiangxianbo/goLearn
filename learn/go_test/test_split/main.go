package main

import (
	"fmt"
	"goLear/go_test/split_string"
)

func main() {
	s := split_string.Split("babcbefb", "b")
	fmt.Printf("%#v", s)
}
