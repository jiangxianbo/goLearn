package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周玲",
		Age:  100,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
	}
	fmt.Printf("%#v\n", string(b))
	// 反序列化
	str := `{"name":"周玲","age":100}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了在函数Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
