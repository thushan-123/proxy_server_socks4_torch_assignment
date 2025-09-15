package main

import (
	"fmt"
	"go-proxy-server/configuration"
	"net"
)

func handleConnection(user net.Conn, config configuration.Socket) {

	userAddress := user.RemoteAddr().(*net.TCPAddr).IP.String()

	fmt.Println("new connection %s", userAddress)

}
