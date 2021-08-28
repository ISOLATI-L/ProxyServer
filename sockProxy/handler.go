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

	var buffer []byte
	_, err := client.Read(buffer)
	if err != nil {
		return err
	}

	client.Write([]byte{0x05, 0x00})
	n, err := client.Read(buffer)
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
		host = net.IP{
			buffer[4],
			buffer[5],
			buffer[6],
			buffer[7],
			buffer[8],
			buffer[9],
			buffer[10],
			buffer[11],
			buffer[12],
			buffer[13],
			buffer[14],
			buffer[15],
			buffer[16],
			buffer[17],
			buffer[18],
			buffer[19],
		}.String()
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
