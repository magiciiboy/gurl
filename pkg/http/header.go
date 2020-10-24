package http

import (
	"fmt"
	"strings"
)

// Header type contains all HTTP headers
// Including Request Headers, General Headers, Entity Headers
type Header map[string][]string

// Add appends a new value to a header
func (h Header) Add(k, v string) {
	ck := CanonicalKey(k)
	h[ck] = append(h[ck], v)
}

// AddLineString appends a new value to a header with
// the value is extracted from a raw string
// For example, "Content-Length: 123"
func (h Header) AddLineString(line string) (key string, err error) {
	elems := strings.Split(line, ":")
	if len(elems) != 2 {
		return "", fmt.Errorf("Invalid header line: `%s`", line)
	}
	key, value := strings.TrimSpace(elems[0]), strings.TrimSpace(elems[1])
	h.Add(key, value)
	return key, nil
}

// Set changes the whole value of a key
func (h Header) Set(k string, v []string) {
	h[CanonicalKey(k)] = v
}

// SetOne does same as Set with only one value
func (h Header) SetOne(k, v string) {
	h[CanonicalKey(k)] = []string{v}
}

// Has checks if the header has a specific key
func (h Header) Has(k string) bool {
	_, ok := h[CanonicalKey(k)]
	return ok
}

// Get returns the first value of a key
func (h Header) Get(k string) string {
	if h.Has(k) {
		v := h[CanonicalKey(k)]
		if len(v) == 0 {
			return ""
		}
		return v[0]
	}
	return ""
}

// GetAll return all values of a header
func (h Header) GetAll(k string) []string {
	if h.Has(k) {
		return h[CanonicalKey(k)]
	}
	return nil
}

// Delete deletes a header
func (h Header) Delete(k string) {
	if h.Has(k) {
		delete(h, CanonicalKey(k))
	}
}

const diffUpper = 'a' - 'A'

// CanonicalKey transform a header key in to
// Canonical form. For example: content-type to Content-Type
func CanonicalKey(k string) string {
	kbyte := []byte(k)
	upper := true

	for i, c := range kbyte {
		if upper && 'a' <= c && c <= 'z' {
			c -= diffUpper
		}
		if !upper && 'A' <= c && c <= 'Z' {
			c += diffUpper
		}
		kbyte[i] = c
		upper = c == '-'
	}
	return string(kbyte)
}
