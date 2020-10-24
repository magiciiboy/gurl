package http

import (
	urllib "github.com/magiciiboy/gurl/pkg/url"
)

// Request is HTTP Request struct
// Follow structure definition at:
// RFC 7230: https://tools.ietf.org/html/rfc7230
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type Request struct {
	Method       string
	URL          *urllib.URL
	IsSecure     bool
	ProtoVersion *ProtoVersion
	Headers      Header
	QueryParams  map[string]string
	Body         string
	Raw          string
}

// DefaultUserAgent of this library
const defaultUserAgent = "gurl-http-client/1.1"

// CreateGETRequest creates a new GET Request
func CreateGETRequest(url string) (*Request, error) {
	URL, err := urllib.ParseURL(url)
	if err == nil {
		msg := "GET " + URL.Path + " HTTP/1.1\r\n" +
			"Host: " + URL.Host + "\r\n" +
			"User-Agent: gurl/0.0.1\r\n" +
			"Accept: */*\r\n\r\n"
		return &Request{
			Method:       "GET",
			URL:          URL,
			Headers:      Header{"User-Agent": []string{defaultUserAgent}},
			ProtoVersion: &HTTPVersion1_1,
			Raw:          msg,
		}, nil
	}
	return nil, err
}

// GetUserAgent returns Content-Type header
func (r *Request) GetUserAgent() string {
	return r.Headers.Get("User-Agent")
}
