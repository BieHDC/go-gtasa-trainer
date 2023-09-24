package main

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"sync"
	"time"
)

// used for debugging once
/*
func makeTempServer() {
	printer := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.RequestURI)
		http.Error(w, "its fine", http.StatusOK)
	}

	http.HandleFunc("/", printer)
	http.ListenAndServe("localhost:8080", nil)
}
*/

var host string = "127.0.0.1"
var port string = ":8080"

func setHost(s string) {
	host = s
}

func getHost() string {
	return host
}

func makeRequest(s string) {
	makeRequestToHost(host, s)
}

func init() {
	http.DefaultClient.Timeout = 500 * time.Millisecond
}

func makeRequestToHost(lhost string, query string) bool {
	//fmt.Println("http://" + lhost + port + query)
	resp, err := http.Get("http://" + lhost + port + query)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	//we only need a simple check for scanning
	return true
}

// low iq implementation, but its fine
func FindInstances() []string {
	var hosters []string
	var norace sync.Mutex
	var wg sync.WaitGroup

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			cidrAddress := ipnet.String()
			p, err := netip.ParsePrefix(cidrAddress)
			if err != nil {
				fmt.Printf("invalid cidr: %s, error %v\n", cidrAddress, err)
				continue
			}
			// 8.8.8.8/24 => 8.8.8.0/24
			p = p.Masked()

			addr := p.Addr()
			for {
				if !p.Contains(addr) {
					break
				}

				wg.Add(1)
				go func(a string) {
					if makeRequestToHost(a, "/hello") {
						//fixme we should check if its really a game
						norace.Lock()
						hosters = append(hosters, a)
						norace.Unlock()
						//log.Println("Found a host:", a)
					}
					wg.Done()
				}(addr.String())

				addr = addr.Next()
			}
		}
	}

	wg.Wait()
	return hosters
}
