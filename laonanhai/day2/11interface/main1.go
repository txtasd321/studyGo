package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫走")
}

// func (c cat) eat() {
// 	fmt.Println("吃鱼")
// }

// !有没有传参很重要，方法必须与接口相同，上例则不行
func (c cat) eat(str string) {
	fmt.Printf("猫吃%s...\n", str)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("jidong")
}

func (c chicken) eat(str string) {
	fmt.Println("吃饲料！")
}

func main() {
	var a1 animal

	bc := cat{
		name: "淘气",
		feet: 4,
	}
	kfc := chicken{
		feet: 2,
	}

	a1 = bc
	a1.eat("鲈鱼")
	fmt.Println(a1)
	fmt.Printf("%T", a1)
	a1 = kfc
	fmt.Println(a1)
}
