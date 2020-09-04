package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func  hello(i int ) {
	fmt.Println("hello goroutine ",i)
	wg.Done()  // wg把计数牌-1
}

func main() {
	wg.Add(5) //wg 计数牌+3
	for i:=0;i<100;i++ {
		go hello(i) //go开启一个goroutine去执行函数
	}
	fmt.Println("hello main")
	//time.Sleep(time.Second)
	wg.Wait()
}
