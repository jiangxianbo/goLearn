package main

import "fmt"

// 空接口

// interface: 关键字
// interface{}: 空接口

// 空接口作为参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "jxb"
	m1["age"] = 1000
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1) // map[age:1000 hobby:[唱 跳 rap] merried:true name:jxb]

	show(false) // type:bool value:false
	show(nil) // type:<nil> value:<nil>
	show(m1) // type:map[string]interface {} value:map[age:1000 hobby:[唱 跳 rap] merried:true name:jxb]
}
