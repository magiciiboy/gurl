package http

import (
	"fmt"
	"strings"
)

// ProtoVersion HTTP uses a "<major>.<minor>" numbering scheme to indicate versions
// of the protocol. This specification defines version "1.1". The
// protocol version as a whole indicates the sender's conformance with
// the set of requirements laid out in that version's corresponding
// specification of HTTP. (See RFC 7230, Section 2.6)
type ProtoVersion struct {
	Major int8
	Minor int8
	Text  string
}

// HTTPVersion1_0 is HTTP/1.0
var HTTPVersion1_0 = ProtoVersion{Major: 1, Minor: 0, Text: "HTTP/1.0"}

// HTTPVersion1_1 is HTTP/1.1
var HTTPVersion1_1 = ProtoVersion{Major: 1, Minor: 1, Text: "HTTP/1.1"}

// HTTPVersion2_0 is HTTP/2.0
var HTTPVersion2_0 = ProtoVersion{Major: 2, Minor: 0, Text: "HTTP/2.0"}

// GetProtolVersionFromText Extract
func GetProtolVersionFromText(text string) (*ProtoVersion, error) {
	parts := strings.Split(text, "/")
	if len(parts) != 2 && parts[0] != "HTTP" {
		return nil, fmt.Errorf("Unrecognized version: %s", text)
	}
	switch parts[1] {
	case "1.0":
		return &HTTPVersion1_0, nil
	case "1.1":
		return &HTTPVersion1_1, nil
	case "2.0":
		return &HTTPVersion2_0, nil
	default:
		return nil, fmt.Errorf("Unrecognized version: %s", text)
	}
}
