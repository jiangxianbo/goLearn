package main

import (
	"fmt"

	"learn/package_test/calc"
)

func init() {
	fmt.Println("自动执行")
}

func main() {
	ret := calc.Add(12, 22)
	fmt.Println(ret)
}
