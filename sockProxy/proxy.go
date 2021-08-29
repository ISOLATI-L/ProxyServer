package sockProxy

import (
	"log"
	"net"
)

func Listen(addr string) {
	// tcpaddr, err := net.ResolveTCPAddr("tcp4", addr)
	// if err != nil {
	// 	log.Fatalln("Error: ", err.Error())
	// }
	// // log.Println(tcpaddr)
	// listener, err := net.ListenTCP("tcp", tcpaddr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for {
		conn, err := listener.Accept()
		log.Printf("Received socket request %s %s\n", conn.LocalAddr().String(), conn.RemoteAddr().String())
		if err != nil {
			log.Fatalln("Error: ", err.Error())
		}

		go func(conn net.Conn) {
			err := handleClientRequest(conn)
			if err != nil {
				log.Println("Error: ", err.Error())
			}
		}(conn)
	}
}
