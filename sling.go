package sling

import (
	"io"
	"net/http"
	"net/url"
)

const (
	contentType     = "Content-Type"
	jsonContentType = "application/json"
	formContentType = "application/x-www-form-urlencoded"
)

// Doer executes http requests.  It is implemented by *http.Client.  You can
// wrap *http.Client with layers of Doers to form a stack of client-side
// middleware.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Sling is an HTTP Request builder and sender.
type Sling struct {
	// http Client for doing requests
	httpClient Doer
	// HTTP method (GET, POST, etc.)
	method string
	// raw url string for requests
	rawURL string
	// stores key-values pairs to add to request's Headers
	header http.Header
	// url tagged query structs
	queryStructs []interface{}
	// body provider
	bodyProvider BodyProvider
	// response decoder
	responseDecoder ResponseDecoder
}

// New returns a new Sling with an http DefaultClient.
func New() *Sling { _ = "STUB: not implemented"; return nil }

// New returns a copy of a Sling for creating a new Sling with properties
// from a parent Sling. For example,
//
//	parentSling := sling.New().Client(client).Base("https://api.io/")
//	fooSling := parentSling.New().Get("foo/")
//	barSling := parentSling.New().Get("bar/")
//
// fooSling and barSling will both use the same client, but send requests to
// https://api.io/foo/ and https://api.io/bar/ respectively.
//
// Note that query and body values are copied so if pointer values are used,
// mutating the original value will mutate the value within the child Sling.
func (s *Sling) New() *Sling {
	_ = "STUB: not implemented"
	// copy Headers pairs into new Header map
	return nil
}

// Http Client

// Client sets the http Client used to do requests. If a nil client is given,
// the http.DefaultClient will be used.
func (s *Sling) Client(httpClient *http.Client) *Sling { _ = "STUB: not implemented"; return nil }

// Doer sets the custom Doer implementation used to do requests.
// If a nil client is given, the http.DefaultClient will be used.
func (s *Sling) Doer(doer Doer) *Sling { _ = "STUB: not implemented"; return nil }

// Method

// Head sets the Sling method to HEAD and sets the given pathURL.
func (s *Sling) Head(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Get sets the Sling method to GET and sets the given pathURL.
func (s *Sling) Get(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Post sets the Sling method to POST and sets the given pathURL.
func (s *Sling) Post(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Put sets the Sling method to PUT and sets the given pathURL.
func (s *Sling) Put(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Patch sets the Sling method to PATCH and sets the given pathURL.
func (s *Sling) Patch(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Delete sets the Sling method to DELETE and sets the given pathURL.
func (s *Sling) Delete(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Options sets the Sling method to OPTIONS and sets the given pathURL.
func (s *Sling) Options(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Trace sets the Sling method to TRACE and sets the given pathURL.
func (s *Sling) Trace(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Connect sets the Sling method to CONNECT and sets the given pathURL.
func (s *Sling) Connect(pathURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Header

// Add adds the key, value pair in Headers, appending values for existing keys
// to the key's values. Header keys are canonicalized.
func (s *Sling) Add(key, value string) *Sling { _ = "STUB: not implemented"; return nil }

// Set sets the key, value pair in Headers, replacing existing values
// associated with key. Header keys are canonicalized.
func (s *Sling) Set(key, value string) *Sling { _ = "STUB: not implemented"; return nil }

// SetBasicAuth sets the Authorization header to use HTTP Basic Authentication
// with the provided username and password. With HTTP Basic Authentication
// the provided username and password are not encrypted.
func (s *Sling) SetBasicAuth(username, password string) *Sling {
	_ = "STUB: not implemented"
	return nil
}

// basicAuth returns the base64 encoded username:password for basic auth copied
// from net/http.
func basicAuth(username, password string) string { _ = "STUB: not implemented"; return "" }

// Url

// Base sets the rawURL. If you intend to extend the url with Path,
// baseUrl should be specified with a trailing slash.
func (s *Sling) Base(rawURL string) *Sling { _ = "STUB: not implemented"; return nil }

// Path extends the rawURL with the given path by resolving the reference to
// an absolute URL. If parsing errors occur, the rawURL is left unmodified.
func (s *Sling) Path(path string) *Sling { _ = "STUB: not implemented"; return nil }

// QueryStruct appends the queryStruct to the Sling's queryStructs. The value
// pointed to by each queryStruct will be encoded as url query parameters on
// new requests (see Request()).
// The queryStruct argument should be a pointer to a url tagged struct. See
// https://godoc.org/github.com/google/go-querystring/query for details.
func (s *Sling) QueryStruct(queryStruct interface{}) *Sling { _ = "STUB: not implemented"; return nil }

// Body

// Body sets the Sling's body. The body value will be set as the Body on new
// requests (see Request()).
// If the provided body is also an io.Closer, the request Body will be closed
// by http.Client methods.
func (s *Sling) Body(body io.Reader) *Sling { _ = "STUB: not implemented"; return nil }

// BodyProvider sets the Sling's body provider.
func (s *Sling) BodyProvider(body BodyProvider) *Sling { _ = "STUB: not implemented"; return nil }

// BodyJSON sets the Sling's bodyJSON. The value pointed to by the bodyJSON
// will be JSON encoded as the Body on new requests (see Request()).
// The bodyJSON argument should be a pointer to a JSON tagged struct. See
// https://golang.org/pkg/encoding/json/#MarshalIndent for details.
func (s *Sling) BodyJSON(bodyJSON interface{}) *Sling { _ = "STUB: not implemented"; return nil }

// BodyForm sets the Sling's bodyForm. The value pointed to by the bodyForm
// will be url encoded as the Body on new requests (see Request()).
// The bodyForm argument should be a pointer to a url tagged struct. See
// https://godoc.org/github.com/google/go-querystring/query for details.
func (s *Sling) BodyForm(bodyForm interface{}) *Sling { _ = "STUB: not implemented"; return nil }

// Requests

// Request returns a new http.Request created with the Sling properties.
// Returns any errors parsing the rawURL, encoding query structs, encoding
// the body, or creating the http.Request.
func (s *Sling) Request() (*http.Request, error) { _ = "STUB: not implemented"; return nil, nil }

// addQueryStructs parses url tagged query structs using go-querystring to
// encode them to url.Values and format them onto the url.RawQuery. Any
// query parsing or encoding errors are returned.
func addQueryStructs(reqURL *url.URL, queryStructs []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// encodes query structs into a url.Values map and merges maps

// url.Values format to a sorted "url encoded" string, e.g. "key=val&foo=bar"

// addHeaders adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func addHeaders(req *http.Request, header http.Header) { _ = "STUB: not implemented"; return }

// Sending

// ResponseDecoder sets the Sling's response decoder.
func (s *Sling) ResponseDecoder(decoder ResponseDecoder) *Sling {
	_ = "STUB: not implemented"
	return nil
}

// ReceiveSuccess creates a new HTTP request and returns the response. Success
// responses (2XX) are JSON decoded into the value pointed to by successV.
// Any error creating the request, sending it, or decoding a 2XX response
// is returned.
func (s *Sling) ReceiveSuccess(successV interface{}) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Receive creates a new HTTP request and returns the response. Success
// responses (2XX) are JSON decoded into the value pointed to by successV and
// other responses are JSON decoded into the value pointed to by failureV.
// If the status code of response is 204(no content) or the Content-Lenght is 0,
// decoding is skipped. Any error creating the request, sending it, or decoding
// the response is returned.
// Receive is shorthand for calling Request and Do.
func (s *Sling) Receive(successV, failureV interface{}) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do sends an HTTP request and returns the response. Success responses (2XX)
// are JSON decoded into the value pointed to by successV and other responses
// are JSON decoded into the value pointed to by failureV.
// If the status code of response is 204(no content) or the Content-Length is 0,
// decoding is skipped. Any error sending the request or decoding the response
// is returned.
func (s *Sling) Do(req *http.Request, successV, failureV interface{}) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// when err is nil, resp contains a non-nil resp.Body which must be closed

// The default HTTP client's Transport may not
// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
// not read to completion and closed.
// See: https://golang.org/pkg/net/http/#Response

// Don't try to decode on 204s or Content-Length is 0

// Decode from json

// decodeResponse decodes response Body into the value pointed to by successV
// if the response is a success (2XX) or into the value pointed to by failureV
// otherwise. If the successV or failureV argument to decode into is nil,
// decoding is skipped.
// Caller is responsible for closing the resp.Body.
func decodeResponse(resp *http.Response, decoder ResponseDecoder, successV, failureV interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
