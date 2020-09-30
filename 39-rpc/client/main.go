package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	rpcdemo "oliver/study/39-rpc"
)

func main() {
	dial, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(dial)

	var result float64
	client.Call("DemoService.Div", rpcdemo.Args{10,3}, &result)
	fmt.Println(result)

	client.Call("DemoService.Div", rpcdemo.Args{4,3}, &result)
	fmt.Println(result)

}
