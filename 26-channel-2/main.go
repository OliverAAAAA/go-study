package main

import "fmt"

func send(ch chan <- int) { //参数： 只读通道

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func rece(ch1, ch2 chan int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go send(ch1)
	go rece(ch1,ch2)
	for ret := range ch2{
		fmt.Println(ret)
	}
}
