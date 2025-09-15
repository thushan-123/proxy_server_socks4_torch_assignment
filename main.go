package main

import (
	"fmt"
	"go-proxy-server/configuration"
	"net"
)

func main() {

	config := configuration.SocketConfigeration()

	listing, err := net.Listen("tcp", config.Address)

	if err != nil {
		fmt.Println("fail listing") // fail to lisning
	}

	defer listing.Close() // close the connection

	println("SOCKS5 proxy run %s", config.Address)

	for {
		connection, err := listing.Accept()

		if err != nil {
			fmt.Println("error accept connection %v", err)
			continue // error occour after re run server
		}

		fmt.Println("accept -> ", connection.RemoteAddr())

		connection.Close()
	}
}
