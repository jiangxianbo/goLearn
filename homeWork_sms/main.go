package main

// 学生管理系统

import (
	"fmt"
	"os"
)

var smr student_mgr // 声明一个全局变量

func showMenu(){
	fmt.Println("\n学生管理系统")
	fmt.Println(`
		1.查看所有
		2.新增
		3.修改
		4.删除
		5.退出
	`)
}

func main() {
	smr = student_mgr{
		allStudent: make(map[int64]student),
	}
	for{
		showMenu()
		fmt.Print("请输入操作:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你的选择是：%d\n", choice)
		switch choice {
		case 1:
			smr.showAllStu()
		case 2:
			smr.addStu()
		case 3:
			smr.editStu()
		case 4:
			smr.delStu()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("选错了~")
		}
	}
}
