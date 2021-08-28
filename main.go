package main

import (
	"proxy/httpProxy"
	"proxy/sock5Proxy"
)

const HTTP_PROXY_PORT uint16 = 9090
const SOCK_PROXY_PORT uint16 = 9091

func main() {
	httpProxy.ProxyHttp(HTTP_PROXY_PORT)
	sock5Proxy.ProxySock(SOCK_PROXY_PORT)
}
