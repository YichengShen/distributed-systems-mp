package tcp

import (
	"../messages"
	"fmt"
	"log"
	"net"
)

type Server struct {
	ln   net.Listener
	quit chan interface{}
}

func ServerProcess() {
	// Create a server
	srv := &Server{
		quit: make(chan interface{}),
	}

	// Server starts to listen
	PORT := ":8080"
	ln, err := net.Listen("tcp", PORT)
	fmt.Println("Server listening on", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	srv.ln = ln
	srv.Serve()
}

// Serve contains the loop for accepting new connections
// It returns when Stop is called
func (srv *Server) Serve() {
	for {
		// Accept connection
		conn, err := srv.ln.Accept()
		if err != nil {
			select {
			// Closing the quit channel in Stop will signal Serve to return
			case <-srv.quit:
				return
			default:
				log.Println("accept error", err)
			}
		} else {
			// Create goroutine to handle connection from client
			go srv.HandleConnection(conn)
		}
	}
}

// After the server establish connection with a client,
// HandleConnection handles the communications with the client
func (srv *Server) HandleConnection(conn net.Conn) {
	// Wait for Email
	msg := &messages.Msg{}
	Decode(conn, msg)
	email := msg.Content

	// Print Email
	messages.PrintEmail(email)

	// Send ACK to client
	SendACK(conn)

	// Wait for ACK
	if GotACK(conn) {
		// Close connection with the client
		conn.Close()
		fmt.Println("Connection closed")
		// Shutdown the server
		srv.Stop()
	}
}

// Stop gracefully shuts down the server
// First, it closes the quit channel
// Then, it closes the listener
// Closing the quit channel can break the accepting loop in Serve
func (srv *Server) Stop() {
	fmt.Println("Server shutdown")
	close(srv.quit)
	srv.ln.Close()
}