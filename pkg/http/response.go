package http

import (
	"strconv"
)

// Response is HTTP Response struct
// Follow structure definition at:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type Response struct {
	ProtoVersion *ProtoVersion
	StatusCode   int16
	StatusText   string
	Headers      Header
	Body         string
	Raw          string
	Request      *Request
}

// CreateEmptyResponse constructs an empty response
func CreateEmptyResponse(req *Request) *Response {
	return &Response{
		Headers: Header{},
		Body:    "",
		Raw:     "",
		Request: req,
	}
}

// GetContentType returns Content-Type header
func (res *Response) GetContentType() string {
	return res.Headers.Get("Content-Type")
}

// GetContentLength returns length of body of a response which
// is extracted from headers
func (res *Response) GetContentLength() int64 {
	v := res.Headers.Get("Content-Length")
	if v != "" {
		l, err := strconv.Atoi(v)
		if err == nil {
			return int64(l)
		}
	}
	return -1
}

// GetTransferEncoding returns Transfer-Encoding header
func (res *Response) GetTransferEncoding() string {
	return res.Headers.Get("Transfer-Encoding")
}

// AppendRaw appends raw content to a response
func (res *Response) AppendRaw(s string, isLine bool) {
	res.Raw += s
	if isLine {
		res.Raw += "\r\n"
	}
}
