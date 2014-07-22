package graphviz

import "fmt"

type GraphVizGraph interface {
	Graph() *GraphBase
}

type Graph interface {
	AddAttribute(*Attr)
	AddNode(GraphVizNode)
	AddSubGraph(GraphVizGraph)
	AddRelation(*Relation)
	GraphViz() string
}

type GraphBase struct {
	Type       string
	ID         string
	Properties Properties
	Attributes []*Attr
	Relations  []*Relation

	SubGraphs []Graph
	Nodes     []*Node

	Parent *Graph
}

func (g *GraphBase) Graph() *GraphBase {
	return g
}

func (g *GraphBase) AddAttribute(a *Attr) {
	g.Attributes = append(g.Attributes, a)
}

func (g *GraphBase) AddSubGraph(sub GraphVizGraph) {
	g.SubGraphs = append(g.SubGraphs, sub.Graph())
}

func (g *GraphBase) AddRelation(r *Relation) {
	g.Relations = append(g.Relations, r)
}

func (g *GraphBase) AddNode(gvn GraphVizNode) {
	n := gvn.Node()
	g.Nodes = append(g.Nodes, n)
	for _, r := range n.Relations {
		g.AddRelation(r)
	}
}

func NewDigraph(name string) *GraphBase {
	return NewGraph("digraph", name)
}

func NewClusterSubGraph(label string) *GraphBase {
	g := NewSubGraph(label)
	g.ID = fmt.Sprintf("clusterSub%p", g)
	return g
}

func NewSubGraph(label string) *GraphBase {
	n := newGraph("subgraph")
	n.ID = fmt.Sprintf("sub%p", n)
	n.Properties["label"] = label
	return n
}

func newGraph(graphType string) *GraphBase {
	g := GraphBase{}
	g.Type = graphType
	g.Attributes = make([]*Attr, 0)
	g.Relations = make([]*Relation, 0)
	g.Properties = make(Properties)
	g.SubGraphs = make([]Graph, 0)
	g.Nodes = make([]*Node, 0)

	return &g
}

func NewGraph(graphType string, id string) *GraphBase {
	g := newGraph(graphType)
	g.ID = id

	return g
}

func (g *GraphBase) GraphViz() string {
	const graphTemplate = `
{{.Type}} {{.ID}} {
{{ range $name, $val := .Properties}} {{$name}}="{{$val}}";
{{end}}
{{ range $i, $attr := .Attributes}} {{$attr.Name}} [{{range $n, $v := $attr.Properties}} {{$n}}="{{$v}}"{{end}} ]; {{end}}
{{ range $i, $s := .SubGraphs}} {{ $s.GraphViz }}
{{end}}
{{ range $i, $n := .Nodes}} {{ $n.GraphViz }}
{{end}}
{{ range $i, $r := .Relations}} {{ $r.GraphViz }}
{{end}}
}
`

	return RenderTemplate(graphTemplate, g)
}
