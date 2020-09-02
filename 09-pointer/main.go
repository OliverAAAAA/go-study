package main

import "fmt"

func main() {

	a := 10
	b := &a
	fmt.Printf("%T\n", a)        //int
	fmt.Printf("%T\n", b)        //*int
	fmt.Printf("%v %p\n", a, &a) //10 0xc00001c090
	fmt.Println(a, &a)
	fmt.Printf("%v %p\n", b, &b) //0xc00001c090 0xc00000e028
	fmt.Println(b, &b)
	c := *b        						//指针取值
	fmt.Println(c) 						//10

	modify1(a)
	fmt.Println(a) 						//10
	modify2(&a)
	fmt.Println(a) 						//100
}

func modify1(a int) {
	a = 100
}
func modify2(b *int) {
	*b = 100
}
