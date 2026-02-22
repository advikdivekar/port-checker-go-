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
	Refused  bool
}

func ScanPort(target string, port int, timeout time.Duration) ScanResult {
	start := time.Now()

	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, timeout)

	duration := time.Since(start)

	if err != nil {
		// Detect timeout
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return ScanResult{Port: port, Open: false, Duration: duration, Timedout: true, Refused: false}
		}
		return ScanResult{Port: port, Open: false, Duration: duration, Timedout: false, Refused: true}
	}

	conn.Close()
	return ScanResult{Port: port, Open: true, Duration: duration, Timedout: false, Refused: false}
}
