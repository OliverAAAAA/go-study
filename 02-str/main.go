package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"
	s1 := "Oliver"
	fmt.Println(len(s))

	fmt.Println(s + s1)
	s2 := fmt.Sprintf("%s - %s", s, s1)
	fmt.Println(s2)

	s3 := "how do u do ?"
	fmt.Println(strings.Split(s3, " "))

	fmt.Println(strings.Contains(s3, "do"))
	fmt.Println(strings.HasPrefix(s3, "how"))
	fmt.Println(strings.HasSuffix(s3, "how"))
	fmt.Println(strings.Index(s3, "do"))
	fmt.Println(strings.LastIndex(s3, "do"))

	s4 := []string{"how", "do", "u", "do", "?"}
	fmt.Println(strings.Join(s4, "+"))

	var c1 byte = 'c'
	var c2 rune = 'C'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T", c1, c2)

	c3 := "hello欧丽菲"
	for i := 0; i < len(c3); i++ {
		fmt.Printf("%c\n", c3[i])
	}

	//rune
	for _, r := range c3 {
		fmt.Printf("%c", r)
	}
	var imgUrl = `<img src="(https://photo.zastatic.com/images/photo/[\d]/%s/[\d].jpg)">`
	var id = "123456"


}
