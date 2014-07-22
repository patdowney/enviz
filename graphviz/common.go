package graphviz

import (
	"bytes"
	"fmt"
	"regexp"
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
	err = t.Execute(&doc, i)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	re := regexp.MustCompile("\n\n+")
	return re.ReplaceAllLiteralString(doc.String(), "\n\n")

}
