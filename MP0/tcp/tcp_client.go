package tcp

import (
	"../messages"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func ClientProcess() {
	// Create a client
	clt := &Client{}

	// Connect to server
	CONNECT := ":8080"
	conn, err := net.Dial("tcp", CONNECT)
	fmt.Println("Client connected to", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	clt.conn = conn

	clt.MainTask()
}

func (clt *Client) MainTask() {
	// Input email
	email := messages.CmdlineEmail()
	msg := messages.Msg{Type: messages.EmailType, Content: email}
	// Send email
	Encode(clt.conn, msg)

	// Wait for ACK from server
	if GotACK(clt.conn) {
		// Client sends ACK upon receiving server's ACK
		SendACK(clt.conn)
		// Close connection with server
		clt.conn.Close()
		fmt.Println("Connection closed")
	}
}
