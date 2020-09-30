package main


import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcdemo "oliver/study/39-rpc"
)

func main() {

	rpc.Register(rpcdemo.DemoService{})
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for{
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("error")
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}
