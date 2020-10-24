package url

import (
	"fmt"
	"regexp"
)

// URL contains all information of an URL
type URL struct {
	Scheme   string
	Host     string // host
	Port     int    // port
	Path     string // path
	Query    string // query string
	Fragment string // fragment after #
	URL      string
}

// URLPattern is RegExp pattern of an URL
const URLPattern = `^((http[s]?):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)?(#[\w\-]+)?$`

// ParseURL creates an URL object from a string
func ParseURL(url string) URL {
	re := regexp.MustCompile(URLPattern)
	matched := re.MatchString(url)

	fmt.Printf("Match: %v\n", matched)
	parts := re.FindAllString(url, -1)
	for _, elem := range parts {
		fmt.Println(elem)
	}
	return URL{}
}

// ParseSimpleURL creates a simple URL with URL string
func ParseSimpleURL(url string) URL {
	return URL{
		Scheme: "https",
		Host:   "linktree.magicii.workers.dev",
		Port:   443,
		Path:   "/links",
		URL:    url,
	}
}
