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
const URLPattern = `^(?:https?:\/\/)?(?:[^@\/\n]+@)?([^:\/\n]+)/([\w\-\_\.\/]*)?`

// ParseURL is a simple version of extracting an URL object from a string
func ParseURL(url string) (*URL, error) {
	re := regexp.MustCompile(URLPattern)
	matched := re.MatchString(url)

	if matched {

		submatchall := re.FindAllStringSubmatch(url, -1)
		host, path := "", ""
		if len(submatchall) >= 1 {
			parts := submatchall[0]
			if l := len(parts); l == 3 {
				host, path = parts[1], "/"+parts[2]
			} else if l == 2 {
				host, path = parts[1], "/"
			}

			if host != "" && path != "" {
				return &URL{
					Scheme: "http",
					Host:   host,
					Path:   path,
					Port:   80,
					URL:    url,
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("Invalid URL. Got: %v", url)
}
