package graphviz

import "fmt"

type GraphVizNode interface {
	Node() *Node
}

type Node struct {
	ID         string
	Properties Properties
	Relations  []*Relation
	Parent     *Graph
}

func (n *Node) Node() *Node {
	return n
}

func (n *Node) GraphViz() string {
	const graphTemplate = `{{.ID}} [{{ range $n, $v := .Properties}} {{$n}}="{{$v}}"{{end}} ];`

	return RenderTemplate(graphTemplate, n)
}

func (n *Node) AddRelation(right *Node) *Relation {
	r := NewRelation(n.ID, right.ID)
	n.Relations = append(n.Relations, r)
	return r
}

func NewNode(name string) *Node {
	n := Node{Properties: make(Properties),
		Relations: make([]*Relation, 0)}

	n.ID = fmt.Sprintf("n%p", &n)

	n.Properties["label"] = name

	return &n
}
