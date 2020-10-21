package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	// "github.com/magiciiboy/gurl/pkg/request"
)

func sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

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
	}

	fmt.Println("gurl", *url, *profile, *verbose, sum(1, 2, 3, 4, 5))

	// req := request.Request{URL: "http://abc.com"}
	// fmt.Println(req)

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
