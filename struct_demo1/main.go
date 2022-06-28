package main

import "fmt"

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main() {
	var p1 person
	p1.name = "你好"
	p1.age = 18
	p1.gender = "男"
	p1.hobby = []string{"篮球", "足球"}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)

	// 匿名结构体：用于临时场景
	var s struct{
		name string
		age int
	}
	s.age = 19
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
