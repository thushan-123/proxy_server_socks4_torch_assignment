package configuration

import "os"

type Socket struct {
	Address  string
	Username string
	Password string
}

func SocketConfigeration() Socket {
	port := os.Getenv("PROXY_PORT")

	if port == "" {
		port = "1080"
	}

	return Socket{
		Address:  "localhost:" + port,
		Username: "admin",
		Password: "password",
	}

}
