package resources

import "github.com/patdowney/graphviz"

func (c *ServiceCatalogue) Graph() *graphviz.GraphBase {
	g := graphviz.NewClusterSubGraph("Services")

	for _, s := range c.serviceIndex {
		g.AddNode(s)
	}

	return g
}
