package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "1000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parseInt failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v  %T\n", ret1, ret1)

	// Atoi: 字符串转换int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v  %T\n", retInt, retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v  %T\n", boolValue, boolValue) // true  bool

	// 从字符串中解析浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v  %T\n", floatValue, floatValue) // fmt.Printf("%#v  %T\n", floatValue, floatValue)

	// 把数字转换成字符串
	i := 97
	//ret2 := string(i)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)

	// 把int转换成字符串
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v  %T\n", ret3, ret3)

	strconv.AppendBool()
}
