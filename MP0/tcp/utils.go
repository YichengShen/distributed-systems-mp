package tcp

import (
	"../messages"
	"encoding/gob"
	"fmt"
	"net"
)

// SendACK sends a Msg struct with Type = AckType
func SendACK(conn net.Conn) {
	msg := messages.Msg{Type: messages.AckType}
	Encode(conn, msg)
}

// GotACK returns true if ACK received
// Otherwise, it returns false
func GotACK(conn net.Conn) bool {
	// Wait for ACK
	msg := &messages.Msg{}
	Decode(conn, msg)
	// Check whether type matches AckType
	if msg.Type == messages.AckType {
		fmt.Println("Got ACK")
		return true
	}
	return false
}

// Encode sends msg through conn
func Encode(conn net.Conn, msg messages.Msg) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(msg)
}

// Decode receives msg from conn and decodes into msg
func Decode(conn net.Conn, msg *messages.Msg) {
	decoder := gob.NewDecoder(conn)
	decoder.Decode(msg)
}