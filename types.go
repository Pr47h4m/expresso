package expresso

import "html/template"

type Text struct {
	Text string
}

type HTML struct {
	HTML string
}

type JSON map[string]interface{}

type XML struct {
	XML interface{}
}

type File struct {
	Path string
}

type Template struct {
	Template   *template.Template
	Properties interface{}
}

type Formatted struct {
	Text
	HTML
	JSON
	XML
	Default interface{}
}
