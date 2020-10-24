package http

import (
	"bufio"
	"strings"
	"testing"
)

type ResponseReaderTestCase struct {
	raw string
}

var responseReaderTestSet = []ResponseReaderTestCase{
	{
		raw: "HTTP/1.1 200 OK\r\n\r\n",
	},
	{
		raw: "HTTP/1.1 200 OK\r\n" +
			"Date: Sat, 24 Oct 2020 16:43:44 GMT\r\n" +
			"Content-Type: application/json;charset=UTF-8\r\n" +
			"Connection: keep-alive\r\n\r\n",
	},
	{
		raw: "HTTP/1.1 200 OK\r\n" +
			"Date: Sat, 24 Oct 2020 16:43:44 GMT\r\n" +
			"Content-Type: application/json;charset=UTF-8\r\n" +
			"Connection: keep-alive\r\n" +
			"Content-Length: 10\r\n" +
			"\r\n" +
			"0123456789",
	},
}

func TestResponseReader(t *testing.T) {
	for _, tc := range responseReaderTestSet {
		r := bufio.NewReader(strings.NewReader(tc.raw))
		rr := ResponseReader{
			Reader: r,
		}
		res, err := rr.ReadResponse(nil)
		if err != nil {
			t.Errorf("ResponseReader Error: %v", err)
		} else if res.Raw != tc.raw {
			t.Errorf("Different Raw. Expected: %v. Got: %v", tc.raw, res.Raw)
		}
	}
}
