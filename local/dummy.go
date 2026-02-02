package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"sync"
)

func main() {
	var number = 10
	var wg sync.WaitGroup

	listeningPorts := make(chan int)
	for range number {
		wg.Add(1)
		port := 10000 + rand.Intn(10000)
		go startServer(port, listeningPorts, &wg)
	}

	go func() {
		for port := range listeningPorts {
			fmt.Printf("Listening Port: %d\n", port)
		}
	}()

	wg.Wait()
}

func startServer(port int, listeningPorts chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("Error listening on port %d: %v\n", port, err)
		return
	}

	listeningPorts <- port
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			continue
		}

		conn.Close()
	}
}
