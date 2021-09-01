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
	responseCache := cache.Get(r)

	if responseCache == nil {
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
		header := response.Header
		statusCode := response.StatusCode
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("Error: ", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		err = response.Body.Close()
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		responseCache = &cache.Cache{
			Header:     header,
			StatusCode: statusCode,
			Body:       body,
		}
		go cache.Save(responseCache)
	}

	h := w.Header()
	for key, value := range responseCache.Header {
		for _, v := range value {
			h.Add(key, v)
		}
	}
	w.WriteHeader(responseCache.StatusCode) // 将host服务器响应转发回客户端
	_, err := w.Write(responseCache.Body)
	if err != nil {
		log.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
