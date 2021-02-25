package tcp

import (
	"log"
	"net"
)

func Connect(ip, port string) net.Conn{
	// Connect to server
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		log.Panic(err)
	}
	return conn
}
