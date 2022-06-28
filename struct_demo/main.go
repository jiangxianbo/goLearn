package main

import "fmt"

type person struct {
	name, gender string
}

type x struct {
	a int8
	b int8
	c int8
}

func f(x person) {
	x.gender = "女"
}

func f2(x *person) {
	x.gender = "女"
}
func main() {
	var p person
	p.name = "123"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)

	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", &p2)
	fmt.Printf("%p\n", p2)

	var p3 = person{
		name:   "nihao",
		gender: "bu",
	}
	fmt.Printf("%#v\n", p3)

	p4 := &person{
		"nihao",
		"bu",
	}
	fmt.Printf("%#v\n", p4)

	m:= x{
		a: 10,
		b: 20,
		c: 30,
	}
	fmt.Printf("%p\n", &m.a)
	fmt.Printf("%p\n", &m.b)
	fmt.Printf("%p\n", &m.c)
}
