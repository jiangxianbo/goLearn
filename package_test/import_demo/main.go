package main

import (
	"fmt"
	"go_learn/package_test/calc"
)

func init() {
	fmt.Println("自动执行")
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
