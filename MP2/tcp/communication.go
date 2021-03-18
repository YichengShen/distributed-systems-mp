package tcp

import (
	"../msg"
	"encoding/gob"
	"net"
)

// Send sends msg through conn
func Send(conn net.Conn, msg msg.Message) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(msg)
}

// Receive receives msg from conn and decodes into msg
func Receive(conn net.Conn) msg.Message {
	msg := &msg.Message{}
	decoder := gob.NewDecoder(conn)
	decoder.Decode(msg)
	return *msg
}