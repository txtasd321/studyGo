package main

import "fmt"

// TODO同一个结构体可以实现多个接口,接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("猫走")
}

func (c *cat) eat(str string) {
	fmt.Printf("猫吃%s...\n", str)
}

func main() {

}
