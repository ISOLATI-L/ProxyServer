package sockProxy

import (
	"net"
	"proxy/transfer"
	"strconv"
)

const BUFFER_SIZE = 1024

func handleClientRequest(client net.Conn) error {
	if client == nil {
		return nil
	}

	var buffer [1024]byte
	_, err := client.Read(buffer[:])
	if err != nil {
		return err
	}

	client.Write([]byte{0x05, 0x00})
	n, err := client.Read(buffer[:])
	if err != nil {
		return err
	}

	var host, port string
	switch buffer[3] {
	case 0x01: //IPv4
		host = net.IPv4(
			buffer[4],
			buffer[5],
			buffer[6],
			buffer[7],
		).String()
	case 0x03: //domain
		host = string(buffer[5 : n-2])
	case 0x04: //IPv6
		host = append(make(net.IP, 0, 16), buffer[4:20]...).String()
	}
	port = strconv.Itoa(int(buffer[n-2])<<8 | int(buffer[n-1]))

	server, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}
	client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	transferor := &transfer.TwoWayTransferor{
		Stream1: server,
		Stream2: client,
	}
	transferor.Start()
	return nil
}
