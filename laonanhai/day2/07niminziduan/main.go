package main

import "fmt"

// TODO匿名字段
// !适合字段比较少也比较简单的场景
// !不常用！

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"zhou",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)

}
