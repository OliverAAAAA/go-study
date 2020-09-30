package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listen, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	log.Printf("rpc server listen on %s", host)
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	dial, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(dial)
	return client, nil
}
