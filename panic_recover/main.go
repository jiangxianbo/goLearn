package main

import "fmt"

func funcA() {
	fmt.Println("a")
}
func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重错误")
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
