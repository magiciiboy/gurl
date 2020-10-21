package common

import "strings"

// GetContentType extracts header "Content-Type" in a request or a response
func GetContentType(headers map[string][]string) string {
	for header, value := range headers {
		if strings.ToLower(strings.TrimSpace(header)) == "content-type" {
			return strings.ToLower(strings.TrimSpace(value))
		}
	}
	return ""
}
