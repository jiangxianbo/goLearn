package main

import "fmt"

func f5(title string, y ...int) int {
	fmt.Println(y) // y是一个int类型的切片
	return 1

}
func main() {
	f5("lixiang", 1, 2, 3, 4, 5, 6)
}
