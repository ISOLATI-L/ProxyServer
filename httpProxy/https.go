package httpProxy

import (
	"net"
	"net/http"
	"proxy/transfer"
	"time"
)

func httpsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	destConn, err := net.DialTimeout("tcp", r.Host, 60*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	transferor := &transfer.TwoWayTransferor{
		Stream1: destConn,
		Stream2: clientConn,
	}
	transferor.Start()
}
