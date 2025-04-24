package main

import (
	"fmt"
	"net"

	"github.com/John-Hejzlar/httpfromtcp/internal/request"
)

func main() {
	// Setup TCP listener on :42069
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Connection accepted from", conn.RemoteAddr())

		req, err := request.RequestFromReader(conn)
		if err != nil {
			fmt.Println("Error reading request:", err)
		} else {
			fmt.Println("Request line:")
			fmt.Printf("- Method: %s\n", req.RequestLine.Method)
			fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
			fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)
			fmt.Println("Headers:")
			for key, value := range req.Headers {
				fmt.Printf("- %s: %s\n", key, value)
			}
			fmt.Println("Body:")
			fmt.Println(string(req.Body))
		}

		fmt.Println("Connection closed")
	}
}
