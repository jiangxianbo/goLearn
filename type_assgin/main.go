package main

import (
	"fmt"
)

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)

	switch t := a.(type){
	case string:
		fmt.Printf("是个字符串 %v\n", t)
	case int:
		fmt.Printf("是个int %v\n", t)
	case int64:
		fmt.Printf("是个int64 %v\n", t)
	case bool:
		fmt.Printf("是个bool %v\n", t)
	}
}

func main() {
	assign(100)
	assign2(true)
}
