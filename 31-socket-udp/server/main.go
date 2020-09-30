package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	fmt.Println("udp persistserver has been started...")
	if err != nil {
		fmt.Printf("listen error,err%v\n", err)
		return
	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read error,err%v\n", err)
			return
		}
		fmt.Println("接受到的数据:", string(buf[:n]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to %v error,err%v\n", addr,err)
			return
		}
	}
}
