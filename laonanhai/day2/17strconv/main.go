package main

import (
	"fmt"
	"strconv"
)

// strconv数字转字符串类型

func main() {
	str := "10000"
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed, err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret, ret)
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v\n", retInt) // 字符串转int

	i := int(65)

	//ret1 := string(i)

	ret2 := fmt.Sprintf("%d", i)

	ret3 := strconv.Itoa(i) // int转str
	//fmt.Printf("%#v", ret1)
	fmt.Printf("%#v\n", ret2)
	fmt.Printf("%#v\n", ret3)

	// 字符串转bool
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 字符串转浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)

}
