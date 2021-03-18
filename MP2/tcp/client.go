package tcp

import (
	"../msg"
	"fmt"
	"log"
	"net"
)

// Returns a new connection of the client with the server
func NewClient(addr, username string) net.Conn{
	// Connect to server
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Panic(err)
	}

	// Send username to Server
	sendInitialMsg(conn, username)

	return conn
}

// Send username to server
func sendInitialMsg(conn net.Conn, username string) {
	initialMsg := msg.Message{
		To: "_SERVER_",
		From: username,
		Content: username,
	}
	Send(conn, initialMsg)
}

// Wait for messages from server
func HandleServerConnection(conn net.Conn, quitChan chan interface{}) {
	for {
		msg := Receive(conn)
		switch {
		// Send failed. Receiver is not connected
		case msg.From == "_SERVER_":
			fmt.Println(msg.Content)
		// Server says EXIT
		case msg.From == "":
			fmt.Println("Server exit. Exit program.")
			conn.Close()
			close(quitChan) // close quit channel to notify main thread
			return
		// Receive message
		default:
			fmt.Println("From "+msg.From+": "+msg.Content)
		}
	}
}