package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步！")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func (c chicken) move() {
	fmt.Println("鸡动！")
}

func (c chicken) eat(food string) {
	fmt.Println("吃%s\n", food)
}

func main() {
	var a1 animal // 定义一个接口类型的变量
	fmt.Printf("%T\n",a1)

	bc := cat{
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.eat("鱼")
	fmt.Println(a1)
	fmt.Printf("%T\n",a1)


	kfc := chicken{
		feet: 4,
	}
	a1 = kfc
	fmt.Printf("%T\n",a1)
}
