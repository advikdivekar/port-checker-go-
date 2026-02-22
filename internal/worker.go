package internal

import (
	"sync"
	"time"
)

func Worker(
	target string,
	ports <-chan int,
	results chan<- ScanResult,
	timeout time.Duration,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for port := range ports {
		result := ScanPort(target, port, timeout)
		results <- result
	}
}
