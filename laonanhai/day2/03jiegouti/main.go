package main

import "fmt"

/*
new和make都是开辟内存的
new是对基本类型int	bool 	string 	struct能使用，返回的是个对应指针
make是对特定类型map slice chin等使用，返回的就是类型本体map等
*/

type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女"
}

func main() {
	var p person
	p.name = "流量"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)

	var p2 = new(person)
	p2.name = "使者"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)

	// key-value初始化
	var p3 = person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	p4 := person{
		"小王子",
		"男",
	}
	fmt.Printf("%v\n", p4)

}
