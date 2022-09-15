package main

import "fmt"

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

func (c *cat) move() {
	fmt.Println("走猫步！")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

//func (c chicken) move() {
//	fmt.Println("鸡动！")
//}
//
//func (c chicken) eat(food string) {
//	fmt.Println("吃%s\n", food)
//}

func main() {
	var a1 animal
	c1 := cat{"tom", 4}
	c2 := &cat{"jim", 4}
	a1 = &c1 // 使用指针接收者 要传指针
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
