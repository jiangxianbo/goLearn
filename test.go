package main

import "fmt"

type person struct {
	name string
	age  int
}

// 构造函数（构造结构体变量的函数）
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 方法
// 接收者是用对应类型的首字母小写
// 指定了接受者之后，只有接收者这个类型的变量才能调用这个方法
func (p *person) dream(str string) {
	fmt.Printf("%s的梦想是%s\n", p.name, str)
}

//func (p person) guonian()  {
//	p.age ++ // 此处的p是p1的副本，改的是副本
//}

// 指针接收者
// 1.需要改变结构体变量的值时要使用指针接收者
// 2.结构体本身比较大，拷贝的内存开销比较大时也要使用指针接收者
// 3.保持一致性：如果有一个方法使用了指针接收者，其他方法统一也是用指针接收者
func (p *person) guonian() {
	p.age++
}

func main() {
	var p1 person
	p1.name = "xxx"
	p1.age = 18

	p2 := person{"xxx", 18}

	// 调用构造函数
	p3 := newPerson("xxx", 18)
	fmt.Println(p1, p2, p3)
	p1.dream("学go")
	p2.dream("有钱")
	fmt.Println(p1)
	p1.guonian()
	fmt.Println(p1)

}
