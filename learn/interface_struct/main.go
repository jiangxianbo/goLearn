package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套
type animal interface { // 嵌套
	mover
	eater
}
type mover interface {
	move()
}

type eater interface {
	eat()
}

type cat struct {
	name string
	feet int8
}

// cat这个结构体，实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步")
}

// cat这个结构体，实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃了%s...\n", food)

}

func main() {

}
