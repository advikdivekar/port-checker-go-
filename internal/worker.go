package internal

import (
	"sync"
	"time"
)

func Worker(
	target string,
	ports <-chan int,
	results chan<- int,
	timeout time.Duration,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for port := range ports {
		if ScanPort(target, port, timeout) {
			results <- port
		}
	}
}
