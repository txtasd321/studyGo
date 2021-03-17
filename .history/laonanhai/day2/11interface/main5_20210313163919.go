package main

import "fmt"

// TODO类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str := a.(string)
	fmt.Println(str)
}

func main() {
	assign(100)
}
