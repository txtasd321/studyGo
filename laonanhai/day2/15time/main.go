package main

import (
	"fmt"
	"time"
)

// TODO获取时间
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	ret := time.Unix(1564803667, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	// 时间间隔
	fmt.Println(time.Second)
	// now +24小时
	fmt.Println(now.Add(24 * time.Hour))
	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 格式化时间,把语言中时间对象，转换成字符串类型的时间
	// 当前时间
	fmt.Println(now.Format("2006-01-02"))
	//
	fmt.Println(now.Format("2006/01/02 03:04:05"))
	fmt.Println(now.Format("03:04:05"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 03:04:05.000")) // 精确到毫秒
	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-08-03")
	if err != nil {
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	// Sleep
	// time.Sleep(100)
	n := 5
	fmt.Println("开始睡")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒钟过去了")

}
