package main

import (
	"fmt"
)

// TODO使用值接收者和指针接收者的区别

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 方法使用值接收者实现了接口的所有方法
// func (c cat) move() {
// 	fmt.Println("猫走")
// }

// func (c cat) eat(str string) {
// 	fmt.Printf("猫吃%s...\n", str)
// }

// 方法使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("猫走")
}

func (c *cat) eat(str string) {
	fmt.Printf("猫吃%s...\n", str)
}

func main() {
	var a1 animal

	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"假老练", 4} // *cat

	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
