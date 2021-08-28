package main

import (
	"fmt"
	"proxy/httpProxy"
	"proxy/sockProxy"
)

const HTTP_PROXY_PORT uint16 = 9090
const SOCK_PROXY_PORT uint16 = 9091

func main() {
	httpProxy.Listen(fmt.Sprintf(":%d", HTTP_PROXY_PORT))
	sockProxy.Listen(fmt.Sprintf(":%d", SOCK_PROXY_PORT))
}
