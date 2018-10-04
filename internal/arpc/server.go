package arpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func Serve(rpcService interface{}, ipAddress string, port string) {
	err := rpc.Register(rpcService)
	if err != nil {
		log.Fatal("Format of service isn't correct. ", err)
	}
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ipAddress+":"+port)
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Println("Serving on port", port)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
