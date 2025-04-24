package response

import (
	"fmt"
	"io"
	"strconv"

	"github.com/John-Hejzlar/httpfromtcp/internal/headers"
)

type StatusCode int

const (
	StatusOK                  StatusCode = 200
	StatusBadRequest          StatusCode = 400
	StatusInternalServerError StatusCode = 500
)

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	var reason string
	switch statusCode {
	case StatusOK:
		reason = "OK"
	case StatusBadRequest:
		reason = "Bad Request"
	case StatusInternalServerError:
		reason = "Internal Server Error"
	}

	if reason != "" {
		_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", statusCode, reason)
		return err
	}

	_, err := fmt.Fprintf(w, "HTTP/1.1 %d\r\n", statusCode)
	return err
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	return headers.Headers{
		"Content-Length": strconv.Itoa(contentLen),
		"Connection":     "close",
		"Content-Type":   "text/plain",
	}
}

func WriteHeaders(w io.Writer, hs headers.Headers) error {
	for key, value := range hs {
		if _, err := fmt.Fprintf(w, "%s: %s\r\n", key, value); err != nil {
			return err
		}
	}
	// End of headers
	if _, err := io.WriteString(w, "\r\n"); err != nil {
		return err
	}
	return nil
}
