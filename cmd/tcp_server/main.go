package main

import (
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("TCP Sync Server running on :9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		conn.Close()

		
	}
}
