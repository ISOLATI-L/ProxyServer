package main

import (
	"fmt"
	"proxy/httpProxy"
	"proxy/sockProxy"
)

const HTTP_PROXY_PORT uint16 = 9090
const SOCK_PROXY_PORT uint16 = 9091

func main() {
	ch := make(chan struct{}, 2)
	go startListenHttp(ch)
	go startListenSock(ch)
	<-ch
	<-ch
}

func startListenHttp(ch chan struct{}) {
	httpProxy.Listen(fmt.Sprintf(":%d", HTTP_PROXY_PORT))
	ch <- struct{}{}
}

func startListenSock(ch chan struct{}) {
	sockProxy.Listen(fmt.Sprintf(":%d", SOCK_PROXY_PORT))
	ch <- struct{}{}
}
