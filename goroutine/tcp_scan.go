package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Second)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {

	hostname := flag.String("h", "192.168.10.3", "主机名")
	startPort := flag.Int("sp", 0, "端口")
	endPort := flag.Int("ep", 100, "端口")
	timeout := flag.Duration("t", time.Second, "超时时间")

	flag.Parse()

	ports := []int{}
	wg := sync.WaitGroup{}

	for port := *startPort; port <= *endPort; port++ {
		fmt.Println(port)
		wg.Add(1)
		go func(p int) {
			if isOpen(*hostname, p, *timeout) {
				ports = append(ports, p)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
}
