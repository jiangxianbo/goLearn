package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["lixiang"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"])
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有")
	} else {
		fmt.Println(score)
	}

	delete(m1, "lixiang")
}
