package httpProxy

import (
	"crypto/tls"
	"log"
	"net/http"
	"proxy/blacklist"
	"proxy/certificate"
)

func Listen(addr string) {
	cert, err := certificate.GenCertificate() //生成https传输需要用到的密钥对
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	server := http.Server{
		Addr:      addr,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler:   &ProxyHandler{},
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
}

type ProxyHandler struct{}

func (ph *ProxyHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf("Received http request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)

	// 检查host是否在黑名单内
	inBlacklist, err := blacklist.Check(r.Host)
	if err != nil {
		log.Printf("An error occurred while connecting to sql: %s\n", err.Error())
		return
	}
	if inBlacklist {
		log.Printf("%s is in blacklist.\n", r.Host)
		return
	}

	if r.Method == http.MethodConnect {
		httpsHandler(w, r) // 响应https连接
	} else {
		httpHandler(w, r) // 处理http请求
	}
}
