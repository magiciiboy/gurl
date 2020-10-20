package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/magiciiboy/gurl/pkg/request"
)

func main() {
	parser := argparse.NewParser("gurl", "HTTP package")
	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL to request"})
	profile := parser.Int("p", "profile", &argparse.Options{Required: true, Help: "Profile n requests"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	req := request.Request{URL: "http://abc.com"}
	fmt.Println("gurl", url, profile)
	fmt.Println(req)
}
