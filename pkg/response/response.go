package response

import (
	"github.com/magiciiboy/gurl/pkg/common"
	"github.com/magiciiboy/gurl/pkg/protocol"
)

// Request is HTTP Response struct
// Follow structure definition at:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type Request struct {
	ProtoVersion *protocol.ProtoVersion
	StatusCode   int16
	StatusText   string
	Headers      common.Header
	Body         string
}
