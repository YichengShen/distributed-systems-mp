package tcp

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"../app"
)

func UnicastSend(msg app.Message) {
	// Connect to server
	conn := Connect(msg.Destination.Ip, msg.Destination.Port)

	// Print current time
	fmt.Printf(">>> Sent '%v' to process %v. System time is %v \n",
		msg.Content, msg.Destination.Id,
		time.Now().Format("15:04:05.000 Jan _2 2006"))

	// Simulate delay
	delayMillisecond := rand.Intn(msg.Source.MaxDelay - msg.Source.MinDelay + 1) +
		msg.Source.MinDelay
	time.Sleep(time.Duration(delayMillisecond) * time.Millisecond)

	// Send message through conn
	Encode(conn, msg)

	// Close connection
	conn.Close()
}

func UnicastReceive(conn net.Conn) {
	// Wait for message
	msg := &app.Message{}
	Decode(conn, msg)

	// Print receive time
	fmt.Printf("<<< Received '%v' from process %v. System time is %v \n",
		msg.Content, msg.Source.Id,
		time.Now().Format("15:04:05.000 Jan _2 2006"))
}