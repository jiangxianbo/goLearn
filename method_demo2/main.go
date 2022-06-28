package main

import "fmt"

// 任意类型添加方法（自定义类型）
// 不能给别的包里的类型添加方法，只能给自己的包里的类型添加
type myInt int

func (i myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.hello()
}
