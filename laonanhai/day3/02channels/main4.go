package main

import "fmt"

// TODO关闭通道

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)

	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	<-ch1
	<-ch1
	x, ok := <-ch1 // 对于一个已经关闭的通道取值是能取到的，只不过ok变为了false
	fmt.Println(x, ok)
}
