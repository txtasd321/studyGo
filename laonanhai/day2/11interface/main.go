package main

import "fmt"

// TODOinterface的实现

type speaker interface {
	speak() // ?只要实现了speak方法的变量都是speaker类型，方法签名
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	fmt.Println(ss)
}
