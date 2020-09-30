package main

import (
	"flag"
	"fmt"
	"oliver/study/go-spider-distribution/rpcsupport"
	"oliver/study/go-spider-distribution/worker"
)

var port = flag.Int("port", 0, "listen port addr")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must have port")
		return
	}
	rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.SpiderService{})
}
