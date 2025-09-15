package proxy

import (
	"fmt"
	"go-proxy-server/configuration"
	"io"
	"log"
	"net"
	"strconv"
)

func HandleConnection(user net.Conn, config configuration.Socket) {

	defer user.Close()

	userAddress := user.RemoteAddr().(*net.TCPAddr).IP.String()
	fmt.Println("new connection from", userAddress)
	log.Println("new connection from", userAddress)

	// socks5 handshake
	if err := HandleHandshake(user); err != nil {
		fmt.Println("handshake failed:", err)
		log.Println("handshake failed:", err)
		return
	}

	// auth
	if err := UserAuthentication(user, config); err != nil {
		fmt.Println("authentication failed:", err)
		log.Println("authentication failed:", err)
		return
	}

	//read request
	targetAddress, targetPort, err := ReadRequest(user)
	if err != nil {
		fmt.Println("request parsing failed:", err)
		log.Println("request parsing failed:", err)
		return
	}

	//connect target server
	target, err := net.Dial("tcp", net.JoinHostPort(targetAddress, strconv.Itoa(targetPort)))
	if err != nil {
		fmt.Printf("failed to connect target server %s:%d: %v\n", targetAddress, targetPort, err)
		log.Printf("failed to connect target server %s:%d: %v\n", targetAddress, targetPort, err)
		SendResponse(user, 0x04, targetAddress, targetPort) // host unreachable
		return
	}
	defer target.Close()

	// send success
	if err := SendResponse(user, 0x00, targetAddress, targetPort); err != nil {
		fmt.Println("failed to send success response:", err)
		log.Println("failed to send success response:", err)
		return
	}

	fmt.Printf("forwarding %s -> %s:%d\n", userAddress, targetAddress, targetPort)
	log.Printf("forwarding %s -> %s:%d\n", userAddress, targetAddress, targetPort)

	//bidirectional copy
	go func() {
		defer user.Close()
		io.Copy(target, user)
	}()
	io.Copy(user, target)

	fmt.Printf("connection closed %s -> %s:%d\n", userAddress, targetAddress, targetPort)
}
