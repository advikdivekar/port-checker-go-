package main

import (
	"flag"
	"fmt"
	"port-scanner-go/internal"
	"sync"
	"time"
)

func main() {
	target := flag.String("target", "", "Target host")
	startPort := flag.Int("start", 1, "Start port")
	endPort := flag.Int("end", 1024, "End port")
	workers := flag.Int("workers", 100, "Number of workers")
	timeoutSec := flag.Float64("timeout", 1.0, "Timeout in seconds")

	flag.Parse()

	if *target == "" {
		fmt.Println("Target required: -target example.com")
		return
	}

	timeout := time.Duration(*timeoutSec * float64(time.Second))

	//checkpoint 1

	fmt.Printf("\nScanning %s\n", *target)
	fmt.Printf("Ports: %d-%d\n", *startPort, *endPort)
	fmt.Printf("Workers: %d\n", *workers)
	fmt.Printf("Timeout: %.2fs\n\n", *timeoutSec)

	startTime := time.Now()

	ports := make(chan int, *workers)
	results := make(chan int)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go internal.Worker(*target, ports, results, timeout, &wg)
	}

	// Feed ports
	go func() {
		for port := *startPort; port <= *endPort; port++ {
			ports <- port
		}
		close(ports)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	var openPorts []int
	for port := range results {
		fmt.Printf("[+] Port %d is OPEN\n", port)
		openPorts = append(openPorts, port)
	}

	elapsed := time.Since(startTime)

	fmt.Println("\nScan Complete")
	fmt.Println("Open Ports:", openPorts)
	fmt.Printf("Time taken: %.2f seconds\n", elapsed.Seconds())
}
