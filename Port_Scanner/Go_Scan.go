package main

import (
	"fmt"
	"net"
	"sort"
)

// worker_pool scan ports using channels to leverage concurrency
func worker_pool(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// Loop runs the worker_pool concurrently
	for i := 0; i < cap(ports); i++ {
		go worker_pool(ports, results)
	}

	go func() {
		for i := 0; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close((ports))
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("Port: %d\t\tOpen\n", port)
	}
}
