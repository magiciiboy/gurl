package response

import "github.com/magiciiboy/gurl/pkg/common"

// Request is HTTP Response struct
// Follow structure definition at:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type Request struct {
	ProtocolVersion string
	StatusCode      int16
	StatusText      string
	Headers         map[string][]string // Request Headers, General Headers, Entity Headers
	Body            string
}

// GetContentType extracts header "Content-Type" of a response
func (req Request) GetContentType() string {
	return common.GetContentType(req.Headers)
}
