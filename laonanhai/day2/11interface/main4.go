package main

import "fmt"

// TODO空接口

// interface:关键字
// interface{}：空接口

func show(a interface{}) {
	fmt.Printf("type:%T, value:%v\n", a, a)
}

func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "小明"
	m1["age"] = 18
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)

}
