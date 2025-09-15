package main

import (
	"fmt"
	"net"
	"os"
)

type Socket struct {
	address  string
	username string
	password string
}

func socketConfigeration() Socket {
	port := os.Getenv("PROXY_PORT")

	if port == "" {
		port = "1080"
	}

	return Socket{
		address:  "localhost:" + port,
		username: "admin",
		password: "password",
	}

}

func main() {

	config := socketConfigeration()

	listing, err := net.Listen("tcp", config.address)

	if err != nil {
		fmt.Println("fail listing") // fail to lisning
	}

	defer listing.Close() // close the connection

	println("SOCKS5 proxy run %s", config.address)

	for {
		_, err := listing.Accept()

		if err != nil {
			fmt.Println("error accept connection %v", err)
			continue // error occour after re run server
		}

		fmt.Println("accept -> ")
	}
}
