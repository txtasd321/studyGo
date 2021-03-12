package main

import "fmt"

// 自定义类型和类型别名

// type后面跟的是类型
type myInt int     // 自定义类型
type yourInt = int //类型别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)

	var y yourInt
	fmt.Printf("%T\n", y)
}
