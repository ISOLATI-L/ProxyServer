package sockProxy

import (
	"log"
	"net"
)

func Listen(addr string) {
	listener, err := net.Listen("tcp", addr) // 监听
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for {
		conn, err := listener.Accept()
		log.Printf("Received socket request %s %s\n", conn.LocalAddr().String(), conn.RemoteAddr().String()) // 接收到tcp连接
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
