package configuration

import "os"

type Socket struct {
	Address  string
	Username string
	Password string
}

func SocketConfigeration() Socket {
	port := os.Getenv("PROXY_PORT")
	host := os.Getenv("HOST")

	if port == "" {
		port = "1080"
	}

	if host == "" {
		host = "localhost"
	}

	return Socket{
		Address:  host + ":" + port,
		Username: "admin",
		Password: "password",
	}

}
