package graphviz

type Relation struct {
	LeftID     string
	RightID    string
	Properties Properties
}

func (r *Relation) GraphViz() string {
	const graphTemplate = `{{.LeftID}} -> {{.RightID}} [{{ range $n, $v :=     .Properties}} {{$n}}="{{$v}}"{{end}}]`

	return RenderTemplate(graphTemplate, r)
}

func NewRelation(leftID string, rightID string) *Relation {
	r := Relation{
		LeftID:     leftID,
		RightID:    rightID,
		Properties: make(Properties)}

	return &r
}

type Graph interface {
	AddAttribute(*Attr)
	AddNode(*Node)
	AddSubGraph(Graph)
	AddRelation(*Relation)
	GraphViz() string
}

type GraphBase struct {
	Type       string
	Name       string
	Properties Properties
	Attributes []*Attr
	Relations  []*Relation

	SubGraphs []Graph
	Nodes     []*Node
}

func (g *GraphBase) AddAttribute(a *Attr) {
	g.Attributes = append(g.Attributes, a)
}

func (g *GraphBase) AddSubGraph(sub Graph) {
	g.SubGraphs = append(g.SubGraphs, sub)
}

func (g *GraphBase) AddRelation(r *Relation) {
	g.Relations = append(g.Relations, r)
}

func (g *GraphBase) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func NewDigraph(name string) *GraphBase {
	return NewGraph("digraph", name)
}

func NewSubGraph(name string) *GraphBase {
	return NewGraph("subgraph", name)
}

func NewGraph(graphType string, name string) *GraphBase {
	g := GraphBase{}
	g.Type = graphType
	g.Name = name
	g.Attributes = make([]*Attr, 0)
	g.Relations = make([]*Relation, 0)
	g.Properties = make(Properties)
	g.SubGraphs = make([]Graph, 0)
	g.Nodes = make([]*Node, 0)

	return &g
}

func (g *GraphBase) GraphViz() string {
	const graphTemplate = `
{{.Type}} {{.Name}} {
{{ range $name, $val := .Properties}} {{$name}}="{{$val}}"
{{end}}
{{ range $i, $attr := .Attributes}} {{$attr.Name}} [{{range $n, $v := $attr.Properties}} {{$n}}={{$v}}{{end}} ]
{{end}}
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
