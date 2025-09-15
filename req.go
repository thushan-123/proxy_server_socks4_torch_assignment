package main

import (
	"encoding/binary"
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

	switch atyp {
	case atypIPv4:
		buf = make([]byte, 4)
		if _, err := io.ReadFull(user, buf); err != nil {
			return "", 0, err
		}
		addr = net.IP(buf).String()
	case atypDomain:
		buf = make([]byte, 1)
		if _, err := io.ReadFull(user, buf); err != nil {
			return "", 0, err
		}
		domainLen := int(buf[0])
		domain := make([]byte, domainLen)
		if _, err := io.ReadFull(user, domain); err != nil {
			return "", 0, err
		}
		addr = string(domain)
	case atypIPv6:
		buf = make([]byte, 16)
		if _, err := io.ReadFull(user, buf); err != nil {
			return "", 0, err
		}
		addr = net.IP(buf).String()
	default:
		return "", 0, errors.New("unsupported address type")
	}

	// read port
	buf = make([]byte, 2)
	if _, err := io.ReadFull(user, buf); err != nil {
		return "", 0, err
	}
	port := int(binary.BigEndian.Uint16(buf))
	
	return addr, port, nil

}
