package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/magiciiboy/gurl/pkg/http"
)

func main() {
	parser := argparse.NewParser("gurl", "A simplest version of cURL written in Go")

	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL to request"})
	profile := parser.Int("p", "profile", &argparse.Options{Required: false, Help: "Profile n requests"})
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Print details"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	} else {
		fmt.Println("gurl", *url, *profile, *verbose)
		req := http.CreateGETRequest(*url)
		fmt.Printf("Request: %v\n", req)
		res, err := http.DefaultClient.SendRequest(req)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Response: %v\n", res)
		}
	}
}
