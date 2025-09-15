package main

import (
	"fmt"
	"go-proxy-server/configuration"
	"go-proxy-server/proxy"
	"log"
	"net"
)

func main() {

	Init()

	config := configuration.SocketConfigeration()

	// listening
	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		fmt.Printf("failed to listen on %s: %v\n", config.Address, err)
		log.Fatalf("failed to listen on %s: %v", config.Address, err)
	}
	defer listener.Close()

	fmt.Printf("SOCKS5 proxy running on %s\n", config.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept connection: %v\n", err)
			log.Printf("failed to accept connection: %v", err)
			continue
		}

		//client connection handle  concurrently
		go proxy.HandleConnection(conn, config)
	}
}
