package main

import "fmt"

// 方法

// 标识符：变量名 函数名 类型名 方法名
// Go语言中如果表示首字母是大写的，就表示对外部可见（暴漏的、共有的）

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 方法作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型别名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s：汪汪汪~", d.name)
}

// 值接收：传拷贝
//func (p person) guonian() {
//	p.age += 1
//}

// 指针接收：传内存地址
func (p *person) zhenguonian() {
	p.age += 1
}

func (p *person) dream() {
	fmt.Println("500w")
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()
	p1 := newPerson("yuanshuai", 18)
	fmt.Println()

	fmt.Println(p1.age)
	//p1.guonian()
	fmt.Println(p1.age)
	p1.zhenguonian()
	fmt.Println(p1.age)
	p1.dream()
}
