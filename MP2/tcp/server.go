package tcp

import (
	"../msg"
	"fmt"
	"log"
	"net"
)

// the Server struct contains a channel for the stop signal
type Server struct {
	Ln   net.Listener
	Quit chan interface{}
	Addr string
	Conns map[string]net.Conn
}

// NewServer returns a new server
// which keeps accepting new connections
// until Stop is called
func NewServer(addr string) *Server {
	// Create a server
	srv := &Server{
		Quit: make(chan interface{}),
		Addr: addr,
		Conns: make(map[string]net.Conn),
	}

	// Server starts to listen
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Panic(err)
	}
	srv.Ln = ln
	go srv.Serve()
	return srv
}

// Serve contains the loop for accepting new connections
// It returns when Stop is called
func (srv *Server) Serve() {
	for {
		// Accept connection
		conn, err := srv.Ln.Accept()
		if err != nil {
			select {
			// Closing the quit channel in Stop will signal Serve to return
			case <-srv.Quit:
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
	// Wait for initial msg from client
	initialMsg := Receive(conn)
	username := initialMsg.From
	fmt.Println(username + " connected")

	// Map username to conn
	srv.Conns[username] = conn

	// Wait for client messages
	for {
		clientMsg := Receive(conn)

		// Check if receiver is connected
		receiverConn, receiverFound := srv.Conns[clientMsg.To]

		switch {
		// if client exits
		case clientMsg.From == "":
			fmt.Println(username + " disconnected")
			delete(srv.Conns, username)
			conn.Close()
			return
		// if receiver is not connected to Server
		case !receiverFound:
			fmt.Println(
				"Send from " + clientMsg.From +
					" to " + clientMsg.To + " failed. " +
					clientMsg.To + " is not connected",
				)
			Send(conn, msg.Message{
				To: username,
				From: "_SERVER_",
				Content: "Send failed. Receiver not connected.",
			})
		// send message to receiver
		default:
			fmt.Println("Send msg from "+clientMsg.From+" to "+clientMsg.To)
			Send(receiverConn, clientMsg)
		}
	}
}

// Stop gracefully shuts down the server
// First, it closes the quit channel
// Then, it closes the listener
// Closing the quit channel can break the accepting loop in Serve
func (srv *Server) Stop() {
	fmt.Println("Server shutdown")
	close(srv.Quit)
	srv.Ln.Close()
}