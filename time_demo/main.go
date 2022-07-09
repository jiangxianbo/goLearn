package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// time.Unix()
	ret := time.Unix(1656853211, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	// 时间间隔
	fmt.Println(time.Second)

	// now + 1小时
	fmt.Println(now.Add(1 * time.Hour))
	// Sub 两个时间间隔
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2022-07-03 12:25:00")
	if err != nil {
		fmt.Printf("Parse time failed, err%v\n", err)
		return
	}
	sub := now.Sub(nextYear)
	fmt.Printf("sub:%v\n", sub)

	// 定时器
	//timer := time.Tick(time.Second)
	//for t := range timer{
	//	fmt.Println(t)
	//}

	// 格式化时间
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006/01/02"))
	fmt.Println(now.Format("2006/01/02 03:04:05 pm"))
	fmt.Println(now.Format("2006/01/02 03:04:05.000"))

	// 按照对应格式解析字符串类型的时间
	timeObj, err := time.Parse("20060102", "20220703")
	if err != nil {
		fmt.Printf("Parse time failed, err%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// Sleep
	n := 5 // int
	fmt.Println("开始Sleep")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5s过去了")
	//time.Sleep(5 * time.Second)
}

// 时区
func f2() {
	now := time.Now() // 获取本地时间
	fmt.Println(now)
	// 明天你的这个时间
	time.Parse("2006-01-02 15:04:05", "2022-07-04 22:50:04")
	// 按照东八区的失去和格式，去解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed, err%v\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-07-04 22:50:04", loc)
	if err != nil {
		fmt.Println("parse time failed, err:%v\n", err)
		return
	}
	timeSub := timeObj.Sub(now)
	timeObjSub := now.Sub(timeObj)
	fmt.Println(timeSub)
	fmt.Println(timeObjSub)
}
func main() {
	f1()
	//f2()

}
