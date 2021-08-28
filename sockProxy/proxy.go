package sockProxy

import (
	"fmt"
	"log"
	"net"
)

func ProxySock(port uint16) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for {
		client, err := listener.Accept()
		log.Printf("Received socket request %s %s\n", client.LocalAddr(), client.RemoteAddr())
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
