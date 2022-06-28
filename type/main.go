package main

import (
	"fmt"
	"strconv"
)

// 自定义类型和类型别名

// type 后面跟的是类型
type myInt int     // 自定义类型
type yourInt = int // 类型别名

func main() {
	var n myInt
	var m yourInt
	n = 100
	m = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	var s1 string = "false"
	var b bool
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("%T%v\n", b, b)

	a := [...]int{
		97: 1,
		98: 2,
		99: 3,
	}
	fmt.Println(a)

}
