package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b ,err := ioutil.ReadFile("32-http/hello.txt")
	if err !=nil {
		fmt.Printf("err:%v\n",err)
	}
	fmt.Fprintln(w,string(b))
}

func main() {
	http.HandleFunc("/api", sayHello)
	err := http.ListenAndServe(":9090",nil)
	if err  != nil{
		fmt.Printf("start error,err : %v",err)
	}

}
