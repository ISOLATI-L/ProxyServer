package sock5Proxy

import (
	"fmt"
	"log"
	"net"
)

func ProxySock(port uint16) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for {
		client, err := l.Accept()
		if err != nil {
			log.Fatalln("Error: ", err.Error())
		}

		go func(client net.Conn) {
			err := handleClientRequest(client)
			if err != nil {
				log.Println("Error: ", err.Error())
			}
		}(client)
	}
}
