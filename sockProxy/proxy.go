package sockProxy

import (
	"log"
	"net"
)

func Listen(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for {
		client, err := listener.Accept()
		log.Printf("Received socket request %s %s\n", client.LocalAddr().String(), client.RemoteAddr().String())
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
