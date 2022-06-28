package main

import "fmt"

const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

func f(arg int) {
	fmt.Println(arg)
}

func main() {
	f(eat | sleep)
}
