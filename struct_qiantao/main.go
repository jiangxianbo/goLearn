package main

import "fmt"

// 嵌套结构体
type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
	workPlace
}

type company struct {
	name string
	add  address
}

func main() {
	p1 := person{
		name: "zhoulin",
		age:  9000,
		address: address{
			province: "黑龙江",
			city:     "牡丹江",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)

}
