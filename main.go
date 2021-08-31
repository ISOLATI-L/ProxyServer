package main

import (
	"fmt"
	"proxy/httpProxy"
	"proxy/sockProxy"
	"sync"
)

const HTTP_PROXY_PORT uint16 = 9090
const SOCK_PROXY_PORT uint16 = 9091

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go startListenHttp(wg)
	go startListenSock(wg)
	wg.Wait()
}

func startListenHttp(wg sync.WaitGroup) {
	httpProxy.Listen(fmt.Sprintf(":%d", HTTP_PROXY_PORT))
	wg.Done()
}

func startListenSock(wg sync.WaitGroup) {
	sockProxy.Listen(fmt.Sprintf(":%d", SOCK_PROXY_PORT))
	wg.Done()
}
