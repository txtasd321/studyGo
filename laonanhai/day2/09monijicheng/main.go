package main

import "fmt"

// TODO结构体模拟其他语言的继承

type animal struct {
	name string
}

// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

//狗类
type dog struct {
	feet uint8
	animal
}

// 给狗实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s:wangwangwang...\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "旺财"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
