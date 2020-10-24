package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	url := "linktree.magicii.workers.dev"
	service := url + ":80"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	req := "GET /links HTTP/1.1\r\n" +
		"Host: linktree.magicii.workers.dev\r\n" +
		"User-Agent: gurl/0.0.1\r\n" +
		"Accept: */*\r\n\r\n"

	_, err = conn.Write([]byte(req))
	checkError(err)

	r := bufio.NewReader(conn)
	len := -1
	for {
		l, _, err := r.ReadLine()
		checkError(err)
		fmt.Println(string(l))
		if line := string(l); line == "" {
			break
		} else {
			len = getContentLength(line)
		}
	}
	fmt.Println(len)

	p := make([]byte, 1)
	rbody := io.LimitReader(r, 461)
	for {
		n, err := rbody.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF")
				break
			}
			checkError(err)
		}
		fmt.Print(string(p[:n]))
	}

	conn.Close()
	os.Exit(0)
}

func getContentLength(line string) int {
	if strings.Index(strings.ToLower(line), "content-length") == 0 {
		fmt.Println("Yeah")
	}
	return 1
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// Message Length
// RFC 7230, Section 3.3.3
