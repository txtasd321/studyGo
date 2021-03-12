package main

import (
	"fmt"
	"os"
)

// TODO学生管理系统

/*
	函数版学生管理系统
	写一个系统能够查看\新增学生\删除学生
*/

var (
	allStudent map[int64]*string //变量声明
)

type student struct {
	name string
	id   int64
}

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}

}
func addStudeng() {

}

func delStudent() {

}

func main() {
	// 1.打印查单
	allStudent = make(map[int64]*string) // 初始化（开辟内存空间）
	for {
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1.查看学生
		2.新增学生
		3.删除学生
		4.退出
	`)
		fmt.Println("请输入操作序号：")

		// 2.等待用户选择
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d", choice)

		// 3.执行对应函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudeng()
		case 3:
			delStudent()
		case 4:
			os.Exit()
		default:
			fmt.Println("无效输入！")

		}
	}
}
