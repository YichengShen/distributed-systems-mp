package tcp

import (
	"../app"
	"encoding/gob"
	"net"
)

// Encode sends msg through conn
func Encode(conn net.Conn, msg app.Message) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(msg)
}

// Decode receives msg from conn and decodes into msg
func Decode(conn net.Conn, msg *app.Message) {
	decoder := gob.NewDecoder(conn)
	decoder.Decode(msg)
}