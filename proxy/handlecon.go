package proxy

import (
	"errors"
	"io"
	"net"
	"slices"
)

const (
	socksVersion       = 0x05
	authMethodNone     = 0x00
	authMethodUserPass = 0x02
)

func HandleHandshake(client net.Conn) error {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(client, buf); err != nil {
		return err
	}

	if buf[0] != socksVersion {
		return errors.New("unsupported SOCKS version")
	}

	nMethods := int(buf[1])
	methods := make([]byte, nMethods)
	if _, err := io.ReadFull(client, methods); err != nil {
		return err
	}

	// check username pwd  is support method
	hasUserPass := slices.Contains(methods, authMethodUserPass)

	// res  authentication method
	if hasUserPass {
		_, err := client.Write([]byte{socksVersion, authMethodUserPass})
		return err
	}

	_, err := client.Write([]byte{socksVersion, 0xFF})
	return err
}
