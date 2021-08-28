package httpProxy

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"proxy/certificate"
)

func ProxyHttp(port uint16) {
	cert, err := certificate.GenCertificate()
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	server := http.Server{
		Addr:      fmt.Sprintf(":%d", port),
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
	log.Printf("Received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)
	if r.Method == http.MethodConnect {
		httpsHandler(w, r)
	} else {
		httpHandler(w, r)
	}
}
