package main

import "fmt"

var ch1 chan int
var ch2 chan []int
var ch3 chan bool

//chan是引用类型，必须make初始化
func main() {
	ch1 := make(chan int, 5)
	//ch2 := make(chan int) 		//无缓冲区通道
	for i := 0; i < 3; i++ {
		ch1 <- i
	}
	fmt.Println(len(ch1),cap(ch1))
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch1)
	}
	close(ch1)

}
