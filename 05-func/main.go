package main

import "fmt"

func main() {
	s := sayHello("oliver", 25)
	fmt.Println(s)

	r1 := intSum(1)
	r2 := intSum(1, 2)
	r3 := intSum(1, 2, 3)
	r4 := intSum(1, 2, 3, 4)
	r5 := intSum(1, 2, 3, 4, 5)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)

	x, y := calc(200, 300)
	fmt.Println(x, y)

	deferFuc()

	r6 := complicate(10, 20, add)
	r7 := complicate(10, 20, sub)
	fmt.Println(r6)
	fmt.Println(r7)

}

//			参数1  参数1类型  参数2  参数2类型      返回值   返回值类型
func sayHello(name string, age int) (ret int) {
	fmt.Println(name+"的年龄是:", age)
	ret = age + 10
	return
}

func intSum(num ...int) (result int) {
	//可变参数类型是切片
	fmt.Printf("%T\n", num)
	for _, a := range num {
		result = result + a
	}
	return
}

//         参数         返回值
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

//defer 倒序执行  一般用于释放资源
func deferFuc() {
	fmt.Println("start...")
	defer fmt.Println("111...")
	defer fmt.Println("222...")
	defer fmt.Println("333...")
	fmt.Println("end...")
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}


func complicate(a, b int, op func(int, int) int) int {
	return op(a, b)
}
