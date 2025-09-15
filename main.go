package main

import "os"

type Socket struct {
	address string
	username string
	password string
}

func socketConfigeration () Socket {
	port := os.Getenv("PROXY_PORT")

	if port == "" { 
		port = "1080"
	}

	return Socket{
		address: ":" + port,
		username: "admin",
		password: "password",
	}

}

func main() {



}