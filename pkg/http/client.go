package http

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// Client is an HTTP Client to send HTTP request
// and receive HTTP response
type Client struct {
	PersistConnection net.Conn
	Timeout           time.Duration
	Status            string
}

// DefaultClient is default HTTP client
var DefaultClient = &Client{}

// GetServiceAddress returns service address
// Default port is 80
// TODO: Need to support https and other port
func (c *Client) GetServiceAddress(url string) string {
	return url + ":80"
}

// GetOrResetConnection init, reset or just return the persist connection
// This works since HTTP/1.1
func (c *Client) GetOrResetConnection(service string) (net.Conn, error) {
	if c.Status != "Connecting" {
		conn, err := net.Dial("tcp", service)
		if err != nil {
			return nil, err
		}
		c.PersistConnection = conn
		c.Status = "Connecting"
	}
	return c.PersistConnection, nil
}

// SendRequest sends a HTTP request
func (c *Client) SendRequest(req *Request) (*Response, error) {
	msg := "GET " + req.URL.Path + " HTTP/1.1\r\n" +
		"Host: " + req.URL.Host + "\r\n" +
		"User-Agent: gurl/0.0.1\r\n" +
		"Accept: */*\r\n\r\n"

	addr := c.GetServiceAddress(req.URL.Host)
	conn, err := c.GetOrResetConnection(addr)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte(msg))
	checkError(err)

	r := bufio.NewReader(conn)
	rr := &ResponseReader{
		Reader: r,
	}
	res, err := rr.ReadResponse(req)

	return res, err
}

// CloseConnection closes persistent connection
func (c *Client) CloseConnection() error {
	return c.PersistConnection.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
