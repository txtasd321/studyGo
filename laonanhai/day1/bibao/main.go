package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()

}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)

}

func f3(f func(int, int), x, y int) func() {
	fmt.Println("this is f3")
	tmp := func() {
		f(x, y)
	}
	return tmp

}

func main() {
	ret := f3(f2, 100, 200)
	fmt.Printf("%T\n", ret)
	f1(ret)

}
