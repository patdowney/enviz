package resources

import (
	"github.com/patdowney/graphviz"
)

func (h *Host) Graph() *graphviz.GraphBase {
	g := graphviz.NewClusterSubGraph(h.HostName)
	for _, s := range h.Services {
		g.AddNode(s)
	}

	return g
}
