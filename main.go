package main

import (
	"proxy/httpProxy"
	"proxy/sockProxy"
)

const HTTP_PROXY_PORT uint16 = 9090
const SOCK_PROXY_PORT uint16 = 9091

func main() {
	httpProxy.ProxyHttp(HTTP_PROXY_PORT)
	sockProxy.ProxySock(SOCK_PROXY_PORT)
}
