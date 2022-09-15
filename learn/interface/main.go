package main

import "fmt"

// 引出接口的实例
type speaker interface {
	speak()
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵~")
}

func (d dog) speak() {
	fmt.Println("汪~")
}

func (p person) speak() {
	fmt.Println("啊~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	ss = d1
	fmt.Println(ss)
}
