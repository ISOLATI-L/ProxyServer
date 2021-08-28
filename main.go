package main

import (
	"proxy/httpProxy"
	"proxy/sockProxy"
)

const HTTP_PROXY_PORT uint16 = 9090
const SOCK_PROXY_PORT uint16 = 9091

func main() {
	httpProxy.Listen(HTTP_PROXY_PORT)
	sockProxy.Listen(SOCK_PROXY_PORT)
}
