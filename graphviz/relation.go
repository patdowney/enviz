package graphviz

type Relation struct {
	LeftID     string
	RightID    string
	Properties Properties
}

func (r *Relation) GraphViz() string {
	const graphTemplate = `{{.LeftID}} -> {{.RightID}} [{{ range $n, $v :=     .Properties}} {{$n}}="{{$v}}"{{end}} ];`

	return RenderTemplate(graphTemplate, r)
}

func NewRelation(leftID string, rightID string) *Relation {
	r := Relation{
		LeftID:     leftID,
		RightID:    rightID,
		Properties: make(Properties)}

	return &r
}
