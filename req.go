package main

import (
	"errors"
	"io"
	"net"
)

const (
	socksVersion = 0x01
	cmd          = 0x01
	atypIPv4     = 0x01
	atypDomain   = 0x03
	atypIPv6     = 0x04
)

func ReadRequest(user net.Conn) (string, int, error) {

	buf := make([]byte, 4)

	if _, err := io.ReadFull(user, buf); err != nil {
		return "", 0, err
	}

	if buf[0] != socksVersion || buf[1] != cmd {
		return "", 0, errors.New("invalid req")
	}

	atyp := buf[3]
	var addr string

	if atyp == atypIPv4 {
		buf = make([]byte, 4)
		if _, err := io.ReadFull(user, buf); err != nil {
			return "", 0, err
		}
		addr = net.IP(buf).String()
	}

	if atyp == atypDomain {
		buf = make([]byte, 1)
		if _, err := io.ReadFull(user, buf); err != nil {
			return "", 0, err
		}
		domainLength := int(buf[0])
		domain := make([]byte, domainLength)

		if _, err := io.ReadFull(user, domain); err != nil {
			return "", 0, err
		}
		addr = string(domain)

	}

}
