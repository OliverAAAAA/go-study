package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen error,err%v\n", err)
		return
	}
	defer conn.Close()
	for  {
		input := bufio.NewReader(os.Stdin)
		s, _ := input.ReadString('\n')
		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("write error,err%v\n", err)
			return
		}

		var  buf [1024]byte
		n,addr,err:= conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read from udp error,err%v\n", err)
			return
		}
		fmt.Printf("read  from  %v value %v\n", addr,n)
	}
}
