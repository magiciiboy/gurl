package http

import "time"

// Client is an HTTP Client to send HTTP request
// and receive HTTP response
type Client struct {
	Timeout time.Duration
}

// DefaultClient is default HTTP client
var DefaultClient = &Client{}

// SendRequest sends a HTTP request
func (c *Client) SendRequest(req *Request) (*Response, error) {
	return nil, nil
}
