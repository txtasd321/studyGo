package main

import "fmt"

// TODO类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	}
	fmt.Println(str)
}

func main() {
	assign(100)
}
