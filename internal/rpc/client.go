package rpc

import (
	"log"
	"net/rpc"
)

func Listen() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}
