package expresso

import (
	"encoding/json"
	"encoding/xml"
	"os"

	"net/http"
	"strconv"
)

type Response struct {
	w http.ResponseWriter
	*Context
}

func responseFromHttpResponseWriter(w http.ResponseWriter) Response {
	return Response{w: w}
}

func (r Response) Send(data interface{}) {
	var bs []byte
	switch data := data.(type) {
	case Text:
		r.w.Header().Add("Content-Type", "text/plain")
		bs = []byte(data.Text)
	case JSON:
		r.w.Header().Add("Content-Type", "application/json")
		bs, _ = json.Marshal(data)
	case HTML:
		r.w.Header().Add("Content-Type", "text/html")
		bs = []byte(data.HTML)
	case File:
		bs, _ = os.ReadFile(data.Path)
	case Template:
		// requires special handling
	case XML:
		r.w.Header().Add("Content-Type", "application/xml")
		bs, _ = xml.Marshal(data)
	}

	r.w.Header().Add("x-powered-by", "Expresso")
	status := r.w.Header().Get("Status")
	if status != "" {
		code, _ := strconv.Atoi(status)
		r.w.WriteHeader(code)
	}

	switch data := data.(type) {
	case Template:
		data.Template.Execute(r.w, data.Properties)
	default:
		r.w.Write(bs)
	}
}

func (r Response) Formatted(req *http.Request, data Formatted) {
	accept := req.Header.Get("Accept")
	switch accept {
	case "text/plain":
		r.Send(data.Text)
	case "text/html":
		r.Send(data.HTML)
	case "application/json":
		r.Send(data.JSON)
	case "application/xml":
		r.Send(data.XML)
	default:
		r.Send(data.Default)
	}
}

func (r Response) Status(code int) Response {
	log := r.Context.Logger()
	log.StatusCode = code
	r.w.Header().Add("Status", strconv.Itoa(code))
	return r
}

func (r Response) SendStatus(code int) {
	log := r.Context.Logger()
	log.StatusCode = code
	r.w.WriteHeader(code)
}

func (r Response) Redirect(url string, status int) {
	log := r.Context.Logger()
	log.StatusCode = status
	r.w.Header().Add("Location", url)
	r.w.WriteHeader(status)
}
