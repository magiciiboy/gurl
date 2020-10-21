package request

import (
	// "errors"
	"github.com/magiciiboy/gurl/pkg/common"
)

// Request is HTTP Request struct
// Follow structure definition at:
// RFC 7230: https://tools.ietf.org/html/rfc7230
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type Request struct {
	Method          string
	URL             string
	IsSecure        bool
	ProtocolVersion string
	Headers         map[string][]string // Request Headers, General Headers, Entity Headers
	QueryParams     map[string]string
	Body            string
}

// GetContentType extracts header "Content-Type" of a request
func (req Request) GetContentType() string {
	return common.GetContentType(req.Headers)
}
