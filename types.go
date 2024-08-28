package expresso

import (
	"html/template"
)

// Text represents plain text content for an HTTP response.
type Text struct {
	Content string // The plain text content to be sent in the response.
}

// HTML represents HTML content for an HTTP response.
type HTML struct {
	Content string // The HTML content to be sent in the response.
}

// JSON represents JSON content for an HTTP response.
// It accepts any Go data structure that can be marshaled into JSON.
type JSON struct {
	Data interface{} // The data to be marshaled into JSON for the response.
}

// XML represents XML content for an HTTP response.
// It accepts any Go data structure that can be marshaled into XML.
type XML struct {
	Data interface{} // The data to be marshaled into XML for the response.
}

type YAML struct {
	Data interface{} // The data to be marshaled into YAML for the response.
}

// File represents a file to be sent as an HTTP response.
type File struct {
	Path        string // The file path to be read and sent in the response.
	ContentType string // The MIME type of the file content.
}

// Template represents a parsed HTML template and the data to be applied to it.
type Template struct {
	Tmpl        *template.Template // The parsed HTML template.
	Data        interface{}        // The data to be injected into the template when rendering.
	ContentType string             // The MIME type for the rendered template, defaults to "text/html".
}

// Formatted represents different formats for an HTTP response.
// The appropriate format is selected based on the client's Accept header.
type Formatted struct {
	Text    *Text       // The plain text content option.
	HTML    *HTML       // The HTML content option.
	JSON    *JSON       // The JSON content option.
	XML     *XML        // The XML content option.
	YAML    *YAML       // The YAML content option.
	Default interface{} // The default content if no Accept header matches.
}
