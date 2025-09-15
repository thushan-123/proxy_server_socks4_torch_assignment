package proxy

import (
	"errors"
	"go-proxy-server/configuration"
	"io"
	"net"
)

func UserAuthentication(user net.Conn, config configuration.Socket) error {

	buf := make([]byte, 2)

	if _, err := io.ReadFull(user, buf); err != nil {
		return err
	}

	if buf[0] != 0x01 {
		return errors.New("auth invalid")
	}

	usernameLength := int(buf[1])
	username := make([]byte, usernameLength)

	if _, err := io.ReadFull(user, username); err != nil {
		return err
	}

	buf = make([]byte, 1)
	if _, err := io.ReadFull(user, buf); err != nil {
		return err
	}

	passwordLength := int(buf[0])
	password := make([]byte, passwordLength)
	if _, err := io.ReadFull(user, password); err != nil {
		return err
	}

	if string(username) != config.Username || string(password) != config.Password {
		user.Write([]byte{0x01, 0x01}) // invlied credentials
		return errors.New("invalied credentials")
	}

	_, err := user.Write([]byte{0x01, 0x00}) // match username and pwd
	return err
}
