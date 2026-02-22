package internal

import (
	"fmt"
	"net"
	"time"
)

type ScanResult struct {
	Port     int
	Open     bool
	Duration time.Duration
	Timedout bool
}

func ScanPort(target string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", target, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}
