package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for{
		reader := bufio.NewReader(conn)
		var buf[128]byte
		n,err := reader.Read(buf[:])
		if err!=nil{
			fmt.Printf("read error,err%v\n", err)
			break
		}
		rece:= string(buf[:n])
		fmt.Println("rece data :",rece)
		conn.Write([]byte("ok"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	fmt.Println("server has been started...")
	if err != nil {
		fmt.Printf("listen error,err%v\n", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept error,err%v\n", err)
			continue
		}
		go process(conn)
	}

}
