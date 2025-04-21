# Custom HTTP Server in Go

This project is a low-level, from-scratch implementation of an HTTP server written in Go. It’s part of my learning journey through the [Boot.dev course: Learn HTTP Protocol in Go](https://www.boot.dev/courses/learn-http-protocol-golang).

## 🌐 About

The goal of this project is to understand how HTTP works under the hood by building a simple but functional HTTP server *without using* Go's standard `net/http` server package for high-level routing or request handling.

Instead, this server:

- Listens for TCP connections on a port
- Parses raw HTTP requests manually
- Responds with proper HTTP-formatted responses
- Handles multiple HTTP methods (GET, POST, etc.)
- Supports headers, status codes, and basic routing

## 🧠 What I'm Learning

- The structure of raw HTTP requests and responses
- Parsing request lines, headers, and bodies
- The difference between HTTP/1.0 and HTTP/1.1
- Building a basic router
- Handling concurrency with Goroutines
- Working directly with TCP sockets in Go

## 🛠 Technologies

- Go (Golang)
- `net` package (TCP networking)
- No external libraries — just standard Go packages

