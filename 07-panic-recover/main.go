package main

import "fmt"

func main() {

	a()
	b()
	c()
}

func a() {
	fmt.Println("a")
}
func b() {
	defer func(){
		err:= recover()
		if err != nil{
			fmt.Println("b occure error")
		}
	}()
	panic("panic in b")
}
func c() {
	fmt.Println("c")
}
