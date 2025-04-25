package server

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/John-Hejzlar/httpfromtcp/internal/request"
	"github.com/John-Hejzlar/httpfromtcp/internal/response"
)

type Handler func(w *response.Writer, req *request.Request)

type Server struct {
	ln      net.Listener
	closed  atomic.Bool // added closed flag for server state
	handler Handler     // added handler field
}

func Serve(port int, handler Handler) (*Server, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := &Server{
		ln:      ln,
		handler: handler, // assign handler
	}
	go s.listen() // start listening for connections
	return s, nil
}

func (s *Server) Close() error {
	s.closed.Store(true)
	if s.ln != nil {
		return s.ln.Close()
	}
	return nil
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

	w := response.NewWriter(conn)
	req, err := request.RequestFromReader(conn)

	if err != nil {
		w.WriteStatusLine(response.StatusCodeBadRequest)
		body := []byte(fmt.Sprintf("Error parsing request: %v", err))
		w.WriteHeaders(response.GetDefaultHeaders(len(body)))
		w.WriteBody(body)
		return
	}

	s.handler(w, req)
}
