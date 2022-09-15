package main

import "fmt"

type student struct {
	id   int64
	name string
}

type student_mgr struct {
	allStudent map[int64]student
}

func (s student_mgr) showAllStu() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d，名字：%s\n", stu.id, stu.name)
	}
}
func (s student_mgr) addStu() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	newStu := student{
		id:   id,
		name: name,
	}
	s.allStudent[newStu.id] = newStu

}
func (s student_mgr) editStu() {
	var id int64
	fmt.Print("请输入要修改的学号：")
	fmt.Scanln(&id)
	nowStu, ok := s.allStudent[id]
	if !ok {
		fmt.Print("没有此人")
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d 姓名：%s\n", nowStu.id, nowStu.name)
	var newName string
	fmt.Print("请输入要修改的姓名：")
	fmt.Scanln(&newName)
	nowStu.name = newName
	s.allStudent[id] = nowStu
}
func (s student_mgr) delStu() {
	var id int64
	fmt.Print("请输入要删除的学号：")
	fmt.Scanln(&id)
	_, ok := s.allStudent[id]
	if !ok {
		fmt.Print("没有此人")
	}
	delete(s.allStudent, id)
	fmt.Println("删除成功！")
}
