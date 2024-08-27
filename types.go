package expresso

import "html/template"

// Text represents plain text content for an HTTP response.
// It holds the text to be sent as a response.
type Text struct {
	Text string // The plain text content to be sent in the response.
}

// HTML represents HTML content for an HTTP response.
// It holds the HTML markup to be sent as a response.
type HTML struct {
	HTML string // The HTML content to be sent in the response.
}

// JSON represents JSON content for an HTTP response.
// It is a map that can hold any key-value pairs, where keys are strings and values can be of any type.
type JSON map[string]interface{} // The JSON content to be sent in the response.

// XML represents XML content for an HTTP response.
// It holds a generic interface that can be marshaled into XML.
type XML struct {
	XML interface{} // The XML content to be sent in the response.
}

// File represents a file to be sent as an HTTP response.
// It holds the path to the file that should be read and sent as the response body.
type File struct {
	Path string // The file path to be read and sent in the response.
}

// Template represents a parsed HTML template and the data properties to be applied to it.
// It is used to render HTML templates dynamically.
type Template struct {
	Template   *template.Template // The parsed HTML template.
	Properties interface{}        // The data to be injected into the template when rendering.
}

// Formatted represents a set of different possible formats for an HTTP response.
// It holds fields for plain text, HTML, JSON, XML, and a default response type.
// The appropriate field is selected based on the client's Accept header.
type Formatted struct {
	Text                // The plain text content option.
	HTML                // The HTML content option.
	JSON                // The JSON content option.
	XML                 // The XML content option.
	Default interface{} // The default content to be sent if none of the other formats match.
}
