package tcp

import (
	"fmt"
	"log"
	"net"
)

// the Server struct contains a channel for the stop signal
type Server struct {
	Ln   net.Listener
	Quit chan interface{}
	Addr string
}

// NewServer returns a new server
// which keeps accepting new connections
// until Stop is called
func NewServer(addr string) *Server {
	// Create a server
	srv := &Server{
		Quit: make(chan interface{}),
		Addr: addr,
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
	UnicastReceive(conn)
	conn.Close()
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
