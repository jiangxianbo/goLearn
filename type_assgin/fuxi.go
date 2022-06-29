package main

import "fmt"

// 类型断言

func main() {
	var a interface{} // 定义一个空接口变量a
	a = 100
	// 如何判断a保存的值的具体类型是什么
	// 类型断言
	// x.(T)
	v, ok := a.(int8)
	if ok {
		fmt.Println("猜对了，a是int8", v)
	} else {
		fmt.Println("猜错了，不是int8")
	}

	// 2.swich
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("不对")
	}
}
