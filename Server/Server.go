/*
 * @Author: jjyaoao 
 * @Date: 2023-06-01 09:41:29 
 * @Last Modified by: jjyaoao
 * @Last Modified time: 2023-06-01 09:47:13
 */
package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*reply = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatalf("Format of service Arith isn't correct. %s", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Listener error: %s", err)
	}

	log.Printf("Serving RPC server on port %d", 1234)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %s", err)
			continue
		}
		go func(conn net.Conn) {
			log.Printf("New connection established")
			rpc.ServeConn(conn)
			log.Printf("Connection closed")
		}(conn)
	}
}
