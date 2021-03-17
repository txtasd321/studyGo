package main

import "fmt"

// TODO类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	a.(string)
}

func main() {

}
