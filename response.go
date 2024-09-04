package expresso

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
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
func responseFromHttpResponseWriter(w http.ResponseWriter) Response {
	return Response{w: w}
}

// Send writes the provided data to the HTTP response. It determines the content type
// based on the type of data and sets the appropriate headers. It supports plain text,
// JSON, HTML, XML, and files.
func (r Response) Send(data interface{}) {
	var bs []byte
	var err error

	l := r.Logger()
	switch data := data.(type) {
	case Text:
		r.w.Header().Set("Content-Type", "text/plain")
		bs = []byte(data.Content)
	case *Text:
		r.w.Header().Set("Content-Type", "text/plain")
		bs = []byte(data.Content)
	case JSON:
		r.w.Header().Set("Content-Type", "application/json")
		bs, err = json.Marshal(data.Data)
	case *JSON:
		r.w.Header().Set("Content-Type", "application/json")
		bs, err = json.Marshal(data.Data)
	case HTML:
		r.w.Header().Set("Content-Type", "text/html")
		bs = []byte(data.Content)
	case *HTML:
		r.w.Header().Set("Content-Type", "text/html")
		bs = []byte(data.Content)
	case File:
		bs, err = os.ReadFile(data.Path)
	case Template:
		// Handle Template rendering here (if implemented)
	case XML:
		r.w.Header().Set("Content-Type", "application/xml")
		bs, err = xml.Marshal(data.Data)
	case *XML:
		r.w.Header().Set("Content-Type", "application/xml")
		bs, err = xml.Marshal(data.Data)
	case YAML:
		r.w.Header().Set("Content-Type", "application/x-yaml")
		bs, err = yaml.Marshal(data.Data)
	case *YAML:
		r.w.Header().Set("Content-Type", "application/x-yaml")
		bs, err = yaml.Marshal(data.Data)
	default:
		l.Error("Unsupported response type")
	}

	if err != nil {
		// Log the error and send an internal server error response
		r.Context.Logger().Error(err.Error())
		r.Status(http.StatusInternalServerError).SendStatus(http.StatusInternalServerError)
		return
	}

	r.w.Header().Set("x-powered-by", "Expresso")

	status := r.w.Header().Get("Status")
	if status != "" {
		code, _ := strconv.Atoi(status)
		r.w.WriteHeader(code)
	} else {
		r.w.WriteHeader(http.StatusOK)
	}

	switch data := data.(type) {
	case Template:
		err = data.Tmpl.Execute(r.w, data.Data)
		if err != nil {
			r.Status(http.StatusInternalServerError).SendStatus(http.StatusInternalServerError)
		}
	default:
		r.w.Write(bs)
	}
}

// Formatted sends the response based on the client's Accept header.
func (r Response) Formatted(req *http.Request, data Formatted) {
	accept := strings.Split(req.Header.Get("Accept"), ",")[0]
	switch accept {
	case "text/plain":
		r.Send(data.Text)
	case "text/html":
		r.Send(data.HTML)
	case "application/json":
		r.Send(data.JSON)
	case "application/xml":
		r.Send(data.XML)
	case "application/x-yaml", "text/yaml":
		r.Send(data.YAML)
	default:
		r.Send(data.Default)
	}
}

// Status sets the HTTP status code for the response and logs it.
func (r Response) Status(code int) Response {
	r.Logger().StatusCode = code
	r.w.Header().Set("Status", strconv.Itoa(code))
	return r
}

// SendStatus writes the HTTP status code directly to the response.
func (r Response) SendStatus(code int) {
	r.Logger().StatusCode = code
	r.w.WriteHeader(code)
}

// Redirect sends an HTTP redirect to the specified URL with the given status code.
func (r Response) Redirect(url string, status int) {
	r.Logger().StatusCode = status
	r.w.Header().Set("Location", url)
	r.w.WriteHeader(status)
}
