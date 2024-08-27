package expresso

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Request wraps the standard http.Request with additional fields that make it easier
// to work with HTTP requests in the expresso framework. It includes the parsed request
// body, URL parameters, and query parameters.
type Request struct {
	RawRequest  *http.Request     // The original HTTP request.
	Path        *url.URL          // The URL path of the request.
	Method      string            // The HTTP method used for the request (e.g., GET, POST).
	Headers     http.Header       // The headers included in the request.
	Body        []byte            // The body of the request as a byte slice.
	Params      httprouter.Params // The URL parameters extracted from the request path.
	QueryParams url.Values        // The query parameters parsed from the URL.
}

// requestFromHttpRequest creates a new Request object from an http.Request.
// It reads the request body and parses it into a byte slice, along with other request details.
func requestFromHttpRequest(r *http.Request) *Request {
	// Retrieve the Content-Length header and parse it to determine the size of the request body.
	cLen := r.Header.Get("Content-Length")
	n, _ := strconv.ParseInt(cLen, 0, 64)

	// Allocate a byte slice to hold the request body.
	bs := make([]byte, n)

	// Read the body into the byte slice.
	if _, err := r.Body.Read(bs); err == nil || errors.Is(err, io.EOF) {
		return &Request{
			RawRequest:  r,
			Path:        r.URL,
			Method:      r.Method,
			Headers:     r.Header,
			Body:        bs,
			QueryParams: r.URL.Query(),
		}
	}
	return nil // Return nil if the body could not be read.
}
