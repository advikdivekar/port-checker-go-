package internal

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(target string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", target, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}
