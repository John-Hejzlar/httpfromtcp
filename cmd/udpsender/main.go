package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Resolve the UDP address
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}

	// Dial the UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a new bufio.Reader from os.Stdin
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop for reading from input and sending over UDP
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Println("Error writing to UDP connection:", err)
		}
	}
}
