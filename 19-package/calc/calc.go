package calc

import "fmt"

// 首字母大写，对外可见
var Name = "Oliver"

func init(){
	fmt.Println("这是一个init函数")
}


func Add(a, b int) int {
	return a + b
}
