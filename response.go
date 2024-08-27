package expresso

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"os"
	"strconv"
)

// Response wraps the standard http.ResponseWriter, providing additional methods
// to facilitate sending various types of responses, including JSON, HTML, XML, files,
// and handling response status codes. The Response struct is tightly coupled with
// the Context to log status codes and other response details.
type Response struct {
	w        http.ResponseWriter // The original HTTP response writer.
	*Context                     // The context in which the response is being generated.
}

// responseFromHttpResponseWriter creates a new Response object from an http.ResponseWriter.
// This function initializes a Response with the provided ResponseWriter, allowing further
// processing and response handling.
func responseFromHttpResponseWriter(w http.ResponseWriter) Response {
	return Response{w: w}
}

// Send writes the provided data to the HTTP response. It determines the content type
// based on the type of data and sets the appropriate headers. It supports plain text,
// JSON, HTML, XML, and files.
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

// Formatted sends the response based on the client's Accept header.
// It checks the Accept header in the request and sends the appropriate
// response format (text, HTML, JSON, or XML). If the Accept header does not
// match any of the specified formats, a default response is sent.
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

// Status sets the HTTP status code for the response and logs it.
// The status code is added to the response headers and also logged
// through the associated Context.
func (r Response) Status(code int) Response {
	log := r.Context.Logger()
	log.StatusCode = code
	r.w.Header().Add("Status", strconv.Itoa(code))
	return r
}

// SendStatus writes the HTTP status code directly to the response.
// This method is used to send a response with a specific status code
// without any body content. The status code is also logged.
func (r Response) SendStatus(code int) {
	log := r.Context.Logger()
	log.StatusCode = code
	r.w.WriteHeader(code)
}

// Redirect sends an HTTP redirect to the specified URL with the given status code.
// It sets the "Location" header in the response and writes the status code, effectively
// redirecting the client. The redirect status code is also logged.
func (r Response) Redirect(url string, status int) {
	log := r.Context.Logger()
	log.StatusCode = status
	r.w.Header().Add("Location", url)
	r.w.WriteHeader(status)
}
