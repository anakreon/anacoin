package arpc

import (
	"log"
	"net/rpc"
)

func Connect(ipAddress string, port string) *rpc.Client {
	client, err := rpc.DialHTTP("tcp", ipAddress+":"+port)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}
