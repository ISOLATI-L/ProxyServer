package main

import (
	"proxy/httpProxy"
)

const PROXY_PORT uint16 = 9090

func main() {
	httpProxy.ProxyHttp(PROXY_PORT)
}
