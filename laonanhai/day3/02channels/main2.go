package main

import (
	"fmt"
	"sync"
)

// var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func main() {
	fmt.Println(b)     // nil
	b = make(chan int) // 不带缓冲区的通道初始化，必须使用make进行初始化,
	// b = make(chan int, 16) // 带缓冲区的初始化，为缓冲数量。
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("接收到了b...", x)
	}()

	fmt.Println("向通道中传递19...")
	b <- 19 // 在无缓冲的通道中，若无接收者，程序会报错
	fmt.Println(b)
	wg.Wait()
}
