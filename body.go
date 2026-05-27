package sling

import (
	"io"
)

// BodyProvider provides Body content for http.Request attachment.
type BodyProvider interface {
	// ContentType returns the Content-Type of the body.
	ContentType() string
	// Body returns the io.Reader body.
	Body() (io.Reader, error)
}

// bodyProvider provides the wrapped body value as a Body for reqests.
type bodyProvider struct {
	body io.Reader
}

func (p bodyProvider) ContentType() string { _ = "STUB: not implemented"; return "" }

func (p bodyProvider) Body() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *

	// jsonBodyProvider encodes a JSON tagged struct value as a Body for requests.
	// See https://golang.org/pkg/encoding/json/#MarshalIndent for details.
	new(io.Reader), nil
}

type jsonBodyProvider struct {
	payload interface{}
}

func (p jsonBodyProvider) ContentType() string { _ = "STUB: not implemented"; return "" }

func (p jsonBodyProvider) Body() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// formBodyProvider encodes a url tagged struct value as Body for requests.
// See https://godoc.org/github.com/google/go-querystring/query for details.
type formBodyProvider struct {
	payload interface{}
}

func (p formBodyProvider) ContentType() string { _ = "STUB: not implemented"; return "" }

func (p formBodyProvider) Body() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
