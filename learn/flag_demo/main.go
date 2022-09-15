package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	//name := flag.String("name", "default", "pls input")
	//age := flag.Int("age", 100, "pls input")
	//married := flag.Bool("married", false, "pls input")
	//cTime := flag.Duration("ct", time.Second, "pls input")
	//
	//flag.Parse()
	//fmt.Println(*name)
	//fmt.Println(*age)
	//fmt.Println(*married)
	//fmt.Println(*cTime)

	var name string
	var age int
	var married bool
	var cTime time.Duration
	flag.StringVar(&name, "name", "df", "名字")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "结婚了吗")
	flag.DurationVar(&cTime, "ct", time.Second, "时间")
	flag.Parse()
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(cTime)

	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数
}
