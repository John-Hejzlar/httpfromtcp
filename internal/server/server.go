package server

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/John-Hejzlar/httpfromtcp/internal/response"
)

type Server struct {
	ln     net.Listener
	closed atomic.Bool // added closed flag for server state
}

func Serve(port int) (*Server, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := &Server{
		ln: ln,
	}
	go s.listen() // start listening for connections
	return s, nil
}

func (s *Server) Close() error {
	s.closed.Store(true) // mark server as closed
	return s.ln.Close()
}

func (s *Server) listen() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if s.closed.Load() {
				// Server is closed; exit the loop gracefully.
				return
			}
			// Ignore transient errors.
			continue
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()

	// Write status line
	if err := response.WriteStatusLine(conn, response.StatusOK); err != nil {
		return // ...handle error if needed...
	}
	// Write headers
	hs := response.GetDefaultHeaders(0)
	if err := response.WriteHeaders(conn, hs); err != nil {
		return // ...handle error if needed...
	}
	// Write body
	_, _ = conn.Write([]byte(""))
}
