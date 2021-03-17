package main

import "fmt"

// 空接口

// interface:关键字
// interface{}：空接口
func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "小明"
	m1["age"] = 18
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
}
