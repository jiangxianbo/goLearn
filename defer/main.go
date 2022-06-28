package main

import "fmt"

func deferDemo()  {
	fmt.Println("start")
	defer fmt.Println("1111")
	defer fmt.Println("2222")
	defer fmt.Println("3333")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
