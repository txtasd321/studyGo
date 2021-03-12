package main

import "fmt"

func main() {
	fmt.Print("沙盒")
	fmt.Print("娜扎")
	fmt.Println("沙河")
	fmt.Println("娜扎")

	//获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入内容是：", s)

}
