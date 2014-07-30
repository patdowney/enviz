package main

import (
	"fmt"

	"github.com/patdowney/enviz/resources"
	"github.com/patdowney/graphviz"
)

func graphHosts(hc *resources.HostCatalogue) {
	g := graphviz.NewDigraph("Hosts")
	g.Properties["ranksep"] = "2"
	g.Properties["fontname"] = "Helvetica"
	g.Properties["colorscheme"] = "svg"
	g.Properties["splines"] = "ortho"

	g.AddSubGraph(hc)

	fmt.Printf(g.GraphViz())
}
