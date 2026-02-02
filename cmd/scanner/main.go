package main

import (
	"flag"
	"fmt"
	"os"
	"portScanner/internal/scan"
	"strconv"
	"strings"
)

func main() {
	hostnamePtr := flag.String("h", "localhost", "Target host name")
	portsPtr := flag.String("p", "1-1024", "Port range (e.g. 1-1024)")
	threadsPtr := flag.Int("t", 10, "Number of concurrent threads")

	flag.Parse()
	hostname := *hostnamePtr
	ports := *portsPtr
	threads := *threadsPtr

	startPort, endPort, err := parsePortRange(ports)

	if err != nil {
		fmt.Printf("Error parsing ports: %v\n", err)
		os.Exit(0)
	}

	scanner.Scan(hostname, threads, startPort, endPort)
}

func parsePortRange(rangeStr string) (int, int, error) {
	ports := strings.Split(rangeStr, "-")
	if len(ports) != 2 {
		return 0, 0, fmt.Errorf("invalid formt. use 'start-end'")
	}

	start, err1 := strconv.Atoi(ports[0])
	end, err2 := strconv.Atoi(ports[1])

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("invalid numbers(start: %s, end: %s)", ports[0], ports[1])
	}

	if start > end {
		return 0, 0, fmt.Errorf("start port must be smaller than end port")
	}

	return start, end, nil
}
