package main

import "fmt"

//select 多路复用  ,每次只有一个case满足条件(多个就随机)


func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("nothing")
		}
	}
}