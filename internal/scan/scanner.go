package scanner

import (
	"fmt"
	"net"
	"slices"
	"strconv"
	"sync"
	"time"
)

func Scan(host string, threads int, startPort int, endPort int) {
	var wg sync.WaitGroup

	// チャネルを作る
	ports := make(chan int, 100)
	results := make(chan int, 100)

	// workerの起動
	for range threads {
		wg.Add(1)
		go worker(host, ports, results, &wg)
	}

	// 結果はソートしてから出力するように
	done := make(chan bool)
	var forSort []int
	go func() {
		for p := range results {
			forSort = append(forSort, p)
		}
		done <- true
	}()

	// portsに対象のポートを流し込む
	for i := startPort; i <= endPort; i++ {
		ports <- i
	}
	close(ports)

	wg.Wait()

	close(results)

	<-done

	slices.Sort(forSort)
	for _, p := range forSort {
		fmt.Printf("Port %d is open\n", p)
	}
}

func worker(host string, ports <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		if check_if_port_open(host, p) {
			results <- p
		}
	}
}

func check_if_port_open(host string, port int) bool {
	address := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}
