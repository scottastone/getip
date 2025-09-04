package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
)

func getPublicIP() {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		fmt.Printf("Public IPv4:  not available (no connection)\n")
		return
	}
	ip, _ := io.ReadAll(resp.Body)
	fmt.Printf("Public IPv4: %s\n", ip)
}

func getLocalIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP[len(ipnet.IP)-4:][0] == 172 {
				continue
			}
			if ipnet.IP.To4() != nil {
				fmt.Printf("Local IPv4:  %s\n", ipnet.IP.String())
			}
		}
	}
}

func getIPv6() {
	resp, err := http.Get("https://api6.ipify.org")
	if err != nil {
		fmt.Printf("Public IPv6:  not available\n")
		return
	}
	ip, _ := io.ReadAll(resp.Body)
	fmt.Printf("Public IPv6: %s\n", ip)
}

func main() {
	helpflag := flag.Bool("h", false, "Show this help message.")
	getpublic := flag.Bool("public", true, "Get public IPv4 address")
	getlocal := flag.Bool("local", true, "Get local IPv4 address")
	getipv6 := flag.Bool("ipv6", true, "Get IPv6 address.")
	flag.Parse()

	if *helpflag {
		flag.Usage()
		fmt.Println("Usage: getip [-public] [-local] [-ipv6]")
		return
	}
	if *getpublic {
		getPublicIP()
	}
	if *getlocal {
		getLocalIP()
	}
	if *getipv6 {
		getIPv6()
	}
}
