package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Timeout = time.Duration(2) * time.Second
var LogWG sync.WaitGroup
var PortRange [][]int

func main() {
	ptrTargets := flag.String("t", "", "target range")
	ptrPorts := flag.String("p", "80", "port to scan")
	ptrTimeout := flag.Float64("T", 1, "timeout time in seconds")
	ptrLog := flag.Bool("s", false, "logging to db")
	flag.Parse()

	scan_range := strings.Split(*ptrTargets, ".")
	octet1 := Conv_array_to_int(strings.Split(scan_range[0], "-"))
	octet2 := Conv_array_to_int(strings.Split(scan_range[1], "-"))
	octet3 := Conv_array_to_int(strings.Split(scan_range[2], "-"))
	octet4 := Conv_array_to_int(strings.Split(scan_range[3], "-"))

	if len(octet1) == 1 {
		octet1 = append(octet1, octet1[0])
	}
	if len(octet2) == 1 {
		octet2 = append(octet2, octet2[0])
	}
	if len(octet3) == 1 {
		octet3 = append(octet3, octet3[0])
	}
	if len(octet4) == 1 {
		octet4 = append(octet4, octet4[0])
	}

	ports := strings.Split(*ptrPorts, ",")

	for _, port := range ports {
		if strings.ContainsAny(port, "-") {
			start_port, err_start_range := strconv.Atoi(strings.Split(port, "-")[0])
			end_port, err_end_range := strconv.Atoi(strings.Split(port, "-")[1])

			if err_start_range != nil || err_end_range != nil {
				log.Fatal("port must be a number")
				return
			}
			PortRange = append(PortRange, []int{start_port, end_port, end_port})

		} else {
			port, err := strconv.Atoi(port)

			if err != nil {
				log.Fatal("port must be a number")
				return
			}
			PortRange = append(PortRange, []int{port, port})
		}
	}

	Timeout = time.Duration(*ptrTimeout*1000) * time.Millisecond

	LogWG.Add(1)
	go Log()
	ScanNetworks(octet1, octet2, octet3, octet4)
	LogWG.Wait()
	if *ptrLog {
		SaveOutput()
	}
}
