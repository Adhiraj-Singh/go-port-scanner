package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// attempt to connect to port and just send result to channel
func scanPort(protocol, hostname string, port int, results chan int) {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 500*time.Millisecond)

	if err != nil {
		results <- 0
		return
	}

	conn.Close()
	// intersting formatting i have used here
	results <- port
}

func main() {
	protocol := "tcp"
	hostname := "localhost"

	startPort := 1
	endPort := 1024

	fmt.Printf("Scanning ports %d to %d on %s (concurrent)...\n", startPort, endPort, hostname)

	results := make(chan int)

	totalPorts := endPort - startPort + 1

	for port := startPort; port <= endPort; port++ {
		go scanPort(protocol, hostname, port, results)
	}

	for i := 0; i < totalPorts; i++ {
		port := <-results
		if port != 0 {
			fmt.Printf("Port %d is open!\n", port)
		}
	}

	fmt.Println("All ports scanned!")
}
