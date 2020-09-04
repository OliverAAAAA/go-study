package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial error,err%v\n", err)
		return
	}
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "exit" {
			return
		}
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send error,err%v\n", err)
			return
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("rece error,err%v\n", err)
			return
		}
		fmt.Println("收到服务端:", string(buf[:n]))
	}

}
