package main

import (
	"./msg"
	"./tcp"
	"./utils"
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Parse command-line arguments
	var ip, port, username string
	flag.StringVar(&ip, "ip", "127.0.0.1", "TCP Server IP")
	flag.StringVar(&port, "port", "8080", "TCP Server Port")
	flag.StringVar(&username, "username", "u1", "Username")
	flag.Parse()
	addr := ip + ":" + port

	// Create a new client
	conn := tcp.NewClient(addr, username)

	// Make quit channel
	quitChan := make(chan interface{})

	// Wait for messages from Server
	go tcp.HandleServerConnection(conn, quitChan)

	// Make message channel that waits for user input
	msgChan := make(chan msg.Message)

	// Prompt user
	fmt.Println("\n2 options:")
	fmt.Println("1. Type in the format to send chat message")
	fmt.Println("\tFormat: [receiver's username] [message content]")
	fmt.Println("2. Type 'EXIT' to quit")
	fmt.Print("\n")

	for {
		// Handle command-line input
		go handleUserInput(username, quitChan, msgChan)

		select {
		// quitChan is closed
		// Exit program
		case <-quitChan:
			return
		// Send out the msg that the user entered
		case msg := <-msgChan:
			tcp.Send(conn, msg)
		}
	}
}

// Take command-line input
// If EXIT, close quitChan
// Else, put msg into msgChan
func handleUserInput(username string, quitChan chan interface{}, msgChan chan msg.Message)  {
	// Take user input
	inputCmd := utils.TakeUserInput()
	// Close the quit channel to notify main thread
	if inputCmd == "EXIT" {
		close(quitChan)
	}
	// Parse user input and put msg into msgChan
	inputs := strings.Fields(inputCmd)
	msg := msg.Message{
		To: inputs[0],
		From: username,
		Content: strings.Join(inputs[1:], " "),
	}
	msgChan <- msg
}
