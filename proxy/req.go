package proxy

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
)

const (
	//socksVersion = 0x05
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

func SendResponse(user net.Conn, status byte, addr string, port int) error {

	res := []byte{socksVersion, status, 0x00}

	// check ip
	if ip := net.ParseIP(addr); ip != nil {
		if ip.To4() != nil {
			res = append(res, atypIPv4)
			res = append(res, ip.To4()...)
		} else {
			res = append(res, atypIPv6)
			res = append(res, ip...)
		}
	} else {
		res = append(res, atypDomain)
		res = append(res, byte(len(addr)))
		res = append(res, addr...)
	}

	portBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(portBytes, uint16(port))
	res = append(res, portBytes...)

	_, err := user.Write(res)
	return err
}
