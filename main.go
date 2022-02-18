package main

import (
	"flag"
	"os"
	"strings"
	"time"
)

func UNUSED(x ...interface{}) {}

var Timeout = time.Duration(1) * time.Second
var Multithreading = false
var Verbose = false

func main() {
	args := os.Args
	
	scan_range := strings.Split(args[len(args)-1], ".")
	addr_1 := Conv_array_to_int(strings.Split(scan_range[0], "-"))
	addr_2 := Conv_array_to_int(strings.Split(scan_range[1], "-"))
	subnet := Conv_array_to_int(strings.Split(scan_range[2], "-"))
	host := Conv_array_to_int(strings.Split(scan_range[3], "-"))

	ptrPorts := flag.String("p", "80", "port range")
	ptrTimeout := flag.Float64("t", 1, "timeout time in seconds")
	ptrVerbose := flag.Bool("v", false, "verbosity")
	ptrLog := flag.Bool("log", false, "logging to db")
	flag.Parse()
	
	ports := strings.Split(*ptrPorts, ",")
	Timeout = time.Duration(*ptrTimeout * 1000) * time.Millisecond 
	Verbose = *ptrVerbose

	Scan_ports(addr_1, addr_2, subnet, host, ports)
	//go Err_logger()
	//msg := <- msg_pipe
	//fmt.Println(msg)

	if *ptrLog {
		Save_to_db()
	}
}

