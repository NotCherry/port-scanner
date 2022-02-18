package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup


func Scan_host(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", addr, Timeout)
	if err != nil && Verbose {
		
	}

	if conn != nil {
		tmp := Targets[ip]
		tmp.IP = ip
		tmp.Ports = append(tmp.Ports, Port{Port: port});
		Targets[ip] = tmp
	}
	defer wg.Done()
}

func Scan_hosts(addr_1 int, addr_2 int, subnet int, host []int, port int) {
	if len(host) == 2 {
		for host_i := host[0]; host_i <= host[1]; host_i++ {
			wg.Add(1)
			Scan_host(fmt.Sprintf("%d.%d.%d.%d", addr_1, addr_2, subnet, host_i), port)
		}
	} else {
			wg.Add(1)
			Scan_host(fmt.Sprintf("%d.%d.%d.%d", addr_1, addr_2, subnet, host[0]), port)
	}
}

func Scan_subnet(addr_1 int, addr_2 int, subnet []int, host []int, port int) {
		if len(subnet) == 2 {
			for subnet_i := subnet[0]; subnet_i < subnet[1]; subnet_i++ {
				go Scan_hosts(addr_1, addr_2, subnet_i, host, port)
			}
		} else {
			go Scan_hosts(addr_1, addr_2, subnet[0], host, port)
		}
	}			

func Scan_network(addr_1 []int, addr_2 []int, subnet []int, host []int, port int) {
		if len(addr_1) == 2 {
			for addr_1i := addr_1[0]; addr_1i < addr_1[1]; addr_1i++ {
				if len(addr_2) == 2 {
					for addr_2i := addr_2[0]; addr_2i < addr_2[1]; addr_2i++ {
						Scan_subnet(addr_1i, addr_2i, subnet, host, port)
					} 
				}	else {
					Scan_subnet(addr_1i, addr_2[0], subnet, host, port)
				}
			} 
		}	else {
			if len(addr_2) == 2 {
				for addr_2i := addr_2[0]; addr_2i < addr_2[1]; addr_2i++ {
					Scan_subnet(addr_1[0], addr_2i, subnet, host, port)
				} 
			}	else {
				Scan_subnet(addr_1[0], addr_2[0], subnet, host, port)
			}
		}
}

func Scan_ports(addr_1 []int, addr_2 []int, subnet []int, host []int, ports []string) {
	for _, port := range ports {
		if strings.ContainsAny(port, "-") {
			start_port, err_start_range :=  strconv.Atoi(strings.Split(port, "-")[0])
			end_port, err_end_range := strconv.Atoi(strings.Split(port, "-")[1])
			if err_start_range != nil || err_end_range != nil {
				panic("port must be a number")
			}
			for p := start_port; p <= end_port; p++ {
				Scan_network(addr_1, addr_2, subnet, host, p)
			}
		} else {
			port, err := strconv.Atoi(port)
			if err != nil {
				panic("port must be a number")
			}
			Scan_network(addr_1, addr_2, subnet, host, port)
		}
		wg.Wait()
	}
}