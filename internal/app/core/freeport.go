package core

import (
	"errors"
	"fmt"
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, fmt.Errorf("an error occurred while trying to resolve TCP addr, err: %w", err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, fmt.Errorf("an error occurred while trying to listen TCP on host, err: %w", err)
	}
	defer l.Close()

	v, ok := l.Addr().(*net.TCPAddr)
	if ok {
		return v.Port, nil
	}

	return 0, errors.New("retrieved unexpected net.Addr value")
}
