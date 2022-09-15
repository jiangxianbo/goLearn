package main

import "fmt"

type car interface {
	run()
}

func drive(c car) {
	c.run()
}

func (f falali) run() {
	fmt.Printf("%s速度70迈\n", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度700迈\n", b.brand)
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "法拉利",
	}
	drive(f1)
	drive(b1)
}
