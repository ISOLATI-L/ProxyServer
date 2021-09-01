package httpProxy

import (
	"ProxyServer/cache"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func httpHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	requestCache := cache.GetCache(r)
	if requestCache != nil {
	} else {
	}

	transport := http.DefaultTransport
	request := new(http.Request)
	*request = *r // 复制一份请求，发送给host服务器

	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if prior, ok := request.Header["For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		request.Header.Set("For", clientIP)
	} else {
		log.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	response, err := transport.RoundTrip(request) // 获取host服务器响应
	if err != nil {
		log.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	header := w.Header()
	for key, value := range response.Header {
		for _, v := range value {
			header.Add(key, v)
		}
	}

	w.WriteHeader(response.StatusCode) // 将host服务器响应转发回客户端
	io.Copy(w, response.Body)
	if err != nil {
		log.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	err = response.Body.Close()
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
}
