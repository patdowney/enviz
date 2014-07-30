package resources

import "github.com/patdowney/graphviz"

func (g *HostGroup) Graph() *graphviz.GraphBase {
	subGraph := graphviz.NewClusterSubGraph(g.Name)
	for _, s := range g.Hosts {
		subGraph.AddSubGraph(s)
	}

	return subGraph
}
