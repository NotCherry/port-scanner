package main

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"

	"log"
)

var wg sync.WaitGroup
var totalHosts int32
var foundHosts int32
var logCh = make(chan Target, 1000)

func ScanHost(ip string) {
	var scannedTarget Target
	var found = true

	for _, p := range PortRange {
		for port := p[0]; port <= p[1]; port += 1 {
			addr := fmt.Sprintf("%s:%d", ip, port)
			conn, err := net.DialTimeout("tcp", addr, Timeout)
			if err == nil && conn != nil {

				if found {
					atomic.AddInt32(&foundHosts, 1)
					found = false
				}

				scannedTarget.IP = ip
				scannedTarget.Ports = append(scannedTarget.Ports, Port{Port: port, Filtered: false})

			}
		}
	}

	if scannedTarget.IP != "" {
		Targets.mu.Lock()
		Targets.list[ip] = scannedTarget
		logCh <-scannedTarget
		defer Targets.mu.Unlock()
	}
	defer wg.Done()
}

func ScanNetworks(octet1 []int, octet2 []int, octet3 []int, octet4 []int) {
	for octet1I := octet1[0]; octet1I <= octet2[1]; octet1I++ {
		for octet2I := octet2[0]; octet2I <= octet2[1]; octet2I++ {
			for octet3I := octet3[0]; octet3I <= octet3[1]; octet3I++ {
				for octet4I := octet4[0]; octet4I <= octet4[1]; octet4I++ {
					wg.Add(1)
					go ScanHost(fmt.Sprintf("%d.%d.%d.%d", octet1I, octet2I, octet3I, octet4I))
				}
			}
			wg.Wait()
			totalHosts += foundHosts
			foundHosts = 0
			log.Printf("Found: %d hosts\n", foundHosts)
			log.Printf("Total: %d hosts\n", totalHosts)
		}
	}
	defer LogWG.Done()
}

func Log() {
	for {
		t := <-logCh
		var ports []int
		for p := range t.Ports {
			ports = append(ports, t.Ports[p].Port)
		}
		log.Println(t.IP, ports)
	}
}
