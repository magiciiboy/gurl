package request

import (
	"github.com/magiciiboy/gurl/pkg/common"
	"github.com/magiciiboy/gurl/pkg/protocol"
)

// Request is HTTP Request struct
// Follow structure definition at:
// RFC 7230: https://tools.ietf.org/html/rfc7230
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type Request struct {
	Method       string
	URL          string
	IsSecure     bool
	ProtoVersion *protocol.ProtoVersion
	Headers      common.Header
	QueryParams  map[string]string
	Body         string
}
