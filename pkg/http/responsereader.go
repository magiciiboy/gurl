package http

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// ReadingChunkSize is size of reading chunked content
const ReadingChunkSize = 4

// ResponseReader helps to read/extract http response from IO reader
type ResponseReader struct {
	Reader *bufio.Reader
}

// ReadResponse reads a response from a buffer
func (r *ResponseReader) ReadResponse(req *Request) (res *Response, err error) {
	res = CreateEmptyResponse(req)
	_, err = readFirstLine(r, res)
	if err != nil {
		return
	}

	err = readHeaders(r, res)
	if err != nil {
		return
	}

	_, err = readBodyContent(r, res)
	if err != nil {
		return
	}
	return res, nil
}

// readFirstLine extracts HTTP Protocol version, status code and status text
// For example: HTTP/1.1 200 OK
func readFirstLine(r *ResponseReader, res *Response) (status int, err error) {
	l, _, err := r.Reader.ReadLine()
	if err == nil {
		line := string(l)
		proto, err := GetProtolVersionFromText(line[:8])
		if err == nil {
			status, err := strconv.Atoi(line[9:12])
			if err == nil {
				res.ProtoVersion = proto
				res.StatusCode = int16(status)
				res.StatusText = line[13:]
				res.AppendRaw(line, true)
				return status, nil
			}
		}
	}
	return
}

// readHeaders reads header line-by-line
func readHeaders(r *ResponseReader, res *Response) (err error) {
	for {
		l, _, err := r.Reader.ReadLine()
		if err == nil {
			line := string(l)
			res.AppendRaw(line, true)
			if line == "" {
				break
			} else {
				_, err := res.Headers.AddLineString(line)
				if err != nil {
					break
				}
			}
		}
	}
	return
}

// readBodyContent reads body content based on RFC 7230, Section 3.3.3
func readBodyContent(r *ResponseReader, res *Response) (string, error) {
	// 1. Any response to a HEAD request and any response with a 1xx
	// (Informational), 204 (No Content), or 304 (Not Modified) status
	// code always has empty body
	if (res.StatusCode/100 == 1) || (res.StatusCode == 204) || (res.StatusCode == 304) ||
		(res.Request != nil && res.Request.Method == "HEAD") {
		res.Body = ""
		return "", nil
	}

	contentLen := res.GetContentLength()
	transferEnc := res.GetTransferEncoding()

	body := ""
	var err error
	if transferEnc == "chunked" || contentLen == -1 {
		body, err = readChunkedBodyContent(r, res)
	} else {
		body, err = readLimitedBodyContent(r, res, contentLen)
	}

	res.Body = body
	res.AppendRaw(body, false)
	return body, err
}

func readLimitedBodyContent(r *ResponseReader, res *Response, contentLen int64) (body string, err error) {
	p := make([]byte, ReadingChunkSize)
	rbody := io.LimitReader(r.Reader, 461)
	for {
		n, err := rbody.Read(p)
		if err != nil {
			if err == io.EOF {
				// fmt.Printf("EOF")
				err = nil
				break
			} else {
				return body, err
			}
		}
		body += string(p[:n])
	}
	if int64(len(body)) != contentLen {
		return body, fmt.Errorf("Actual length of message body is different from Content-Length header")
	}
	return body, nil
}

func readChunkedBodyContent(r *ResponseReader, res *Response) (body string, err error) {
	p := make([]byte, ReadingChunkSize)
	rbody := r.Reader
	for {
		n, err := rbody.Read(p)
		if err != nil {
			if err == io.EOF {
				// fmt.Printf("EOF")
				err = nil
				break
			} else {
				return body, err
			}
		}
		body += string(p[:n])
	}
	return body, nil
}
