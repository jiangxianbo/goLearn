package main

import "fmt"

// 结构题遇到的问题
// 1.myInt(100)是什么
type myInt int

type person struct {
	name string
	age  int
}

func main() {
	// 问题1：声明一个myInt类型的变量m，值为100
	//var m myInt
	//m = 100
	//var m myInt = 100
	//var m = myInt(100)
	m := myInt(100)
	fmt.Println(m)

	// 问题2：结构体初始化
	type person struct {
		name string
		age  int
	}
	// 方法1：
	var p person
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "zhoulin"
	p1.age = 1000
	fmt.Println(p1)
	// 方法2：
	var p2 = person{
		name: "官话",
		age:  100,
	}
	fmt.Println(p2)
}

// 问题3:为什么要构造函数
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
