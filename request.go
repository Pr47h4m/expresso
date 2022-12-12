package expresso

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Request struct {
	RawRequest  *http.Request
	Path        *url.URL
	Method      string
	Headers     http.Header
	Body        []byte
	Params      httprouter.Params
	QueryParams url.Values
}

func requestFromHttpRequest(r *http.Request) *Request {
	cLen := r.Header.Get("Content-Length")
	n, _ := strconv.ParseInt(cLen, 0, 64)
	bs := make([]byte, n)
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
	return nil
}
