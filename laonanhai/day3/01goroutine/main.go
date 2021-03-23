package main

import (
	"math/rand"
	"sync"
	"time"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		runtime.Gosched()
// 		fmt.Println(s)
// 	}
// }
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	println(i)
}

func main() {
	// go say("world")
	// say("hello")
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
