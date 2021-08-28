package httpProxy

import (
	"crypto/tls"
	"log"
	"net/http"
	"proxy/certificate"
)

func Listen(addr string) {
	cert, err := certificate.GenCertificate()
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
	if r.Method == http.MethodConnect {
		httpsHandler(w, r)
	} else {
		httpHandler(w, r)
	}
}
