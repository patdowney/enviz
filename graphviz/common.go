package graphviz

import (
	"bytes"
	"fmt"
	"text/template"
)

type Properties map[string]string

type Attr struct {
	Name       string
	Properties Properties
}

func NewAttr(name string) *Attr {
	a := Attr{
		Name:       name,
		Properties: make(Properties)}

	return &a
}

type Graphvizable interface {
	GraphViz() string
}

func RenderTemplate(templateString string, i interface{}) string {
	t, err := template.New("graph").Parse(templateString)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	var doc bytes.Buffer
	t.Execute(&doc, i)
	return doc.String()
}
