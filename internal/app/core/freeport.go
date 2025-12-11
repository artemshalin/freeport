package core

import (
	"errors"
	"fmt"
	"net"
	"time"
)

var ErrNoFreePort = errors.New("no free ports")
var ErrBadRange = errors.New("the 'to' port cannot be smaller than the 'from' port")

func CheckRange(from int, to int, addrPrefix string) (int, error) {
	if from > to {
		return 0, ErrBadRange
	}
	for port := from; port <= to; port++ {
		addr := net.JoinHostPort(addrPrefix, fmt.Sprintf("%d", port))

		conn, err := net.DialTimeout("tcp", addr, 100*time.Millisecond)
		if err != nil {
			return port, nil
		}

		conn.Close()
	}
	return 0, ErrNoFreePort
}
