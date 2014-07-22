package graphviz

type Node struct {
	Name       string
	Properties Properties
}

func (n *Node) GraphViz() string {
	const graphTemplate = `{{.Name}} [{{ range $n, $v := .Properties}} {{$n}}="{{$v}}"{{end}} ];`

	return RenderTemplate(graphTemplate, n)
}

func NewNode(name string) *Node {
	n := Node{Name: name,
		Properties: make(Properties)}
	return &n
}
