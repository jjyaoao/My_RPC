/*
 * @Author: jjyaoao 
 * @Date: 2023-06-01 09:41:26 
 * @Last Modified by:   jjyaoao 
 * @Last Modified time: 2023-06-01 09:41:26 
 */
package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Dialing error: %s", err)
	}

	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatalf("Arith error: %s", err)
	}

	log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
