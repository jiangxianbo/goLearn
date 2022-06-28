package main

import "fmt"

func main() {

	//fmt.Printf("%v\n", 1111)
	//
	//fmt.Printf("%q\n", 65)
	//
	//fmt.Printf("%b\n", 3.14151926123123)
	//fmt.Printf("%q\n", "理想")
	//fmt.Printf("%05s\n", "小王子")
	//fmt.Printf("%5.7s\n", "小王子")



	//var s string
	//fmt.Scan(&s)
	//fmt.Println("用户输入的内容是：", s)

	//var (
	//	name  string
	//	age   string
	//	class string
	//)
	//fmt.Scanf("%s %d %s\n", &name, &age, &class)
	//fmt.Println(name, age, class)

	//fmt.Scanln(&name, &age, &class)
	//fmt.Println(name, age, class)


	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println("-----")
	fmt.Println(s1, s2, s3)
}
