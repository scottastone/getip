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
		panic(err)
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

func main() {
	getpublic := flag.Bool("public", true, "Get public IP address")
	getlocal := flag.Bool("local", true, "Get local IP address")
	flag.Parse()

	if *getpublic {
		getPublicIP()
	}
	if *getlocal {
		getLocalIP()
	}
}
