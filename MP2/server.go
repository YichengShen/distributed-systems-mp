package main

import (
	"./tcp"
	"./utils"
	"flag"
	"fmt"
)

func main() {
	// Parse command-line arguments
	var port string
	flag.StringVar(&port, "port", "8080", "TCP Server Port")
	flag.Parse()

	ip := "127.0.0.1"

	// Launch a server
	// Server waits for messages from other processes
	srv := tcp.NewServer(ip + ":" + port)

	// Prompt user
	fmt.Println("Type 'EXIT' to quit")
	fmt.Print("\n")

	for {
		// Take user input
		inputCmd := utils.TakeUserInput()
		// Break loop if user inputs EXIT
		if inputCmd == "EXIT" {break}
	}

	srv.Stop()
}

