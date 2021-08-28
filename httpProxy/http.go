package httpProxy

import (
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
	transport := http.DefaultTransport
	request := new(http.Request)
	*request = *r

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

	response, err := transport.RoundTrip(request)
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
	// log.Println(w.Header())
	// log.Println(response.Header)
	w.WriteHeader(response.StatusCode)
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
