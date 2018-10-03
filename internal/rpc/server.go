package rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func Serve(rpcService interface{}) {
	err := rpc.Register(rpcService)
	if err != nil {
		log.Fatal("Format of service isn't correct. ", err)
	}
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving on port %d", 1234)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
