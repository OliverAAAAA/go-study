package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10) //wg 计数牌+3

	for i :=0 ;i<10;i++ {
		fmt.Printf("%T\n",i)
		go func(i int ) {
			fmt.Println("hello",i)
			wg.Done()
		}(i) //go开启一个goroutine去执行函数
	}
	fmt.Println("hello main")
	//time.Sleep(time.Second)
	wg.Wait()
}
