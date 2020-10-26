package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/magiciiboy/gurl/pkg/http"
	"github.com/magiciiboy/gurl/pkg/profile"
)

func main() {
	parser := argparse.NewParser("gurl", "A simplest version of cURL written in Go")

	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL to request"})
	prof := parser.Int("p", "profile", &argparse.Options{Required: false, Help: "Profile n requests"})
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Print details"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	} else {
		fmt.Println("gurl", *url, *prof, *verbose)
		req, err := http.CreateGETRequest(*url)
		fmt.Printf("\nRequest:\n%v\n", req.Raw)
		if err != nil {
			fmt.Printf("URL Parsing Error: %s\n", err.Error())
			os.Exit(1)
		}

		if prof == nil || *prof == 0 {
			res, err := http.DefaultClient.SendRequest(req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Response:\n%v\n", res.Raw)
				fmt.Printf("Size: %v B\n", res.GetSize())
			}
		} else {
			p := sendNRequestWithProfiling(http.DefaultClient, req, *prof)
			p.PrintSummary()
		}
		// Close persist connection
		http.DefaultClient.CloseConnection()
	}
}

// sendNRequestWithProfiling sends n requests and profile them
func sendNRequestWithProfiling(c *http.Client, req *http.Request, n int) *profile.SessionProfile {
	p := profile.DefaultProfiler

	for i := 0; i < n; i++ {
		p.DoRequest(c, req)
	}
	p.Profile.DoStat()
	return &p.Profile
}
