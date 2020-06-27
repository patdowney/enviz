package main

import (
	"fmt"

	"github.com/patdowney/enviz/resources"
	"github.com/patdowney/graphviz"
)

func graphHosts(hc *resources.HostCatalogue) {
	g := graphviz.NewDigraph("Hosts")
	g.Properties["ranksep"] = "0.2"
	g.Properties["nodesep"] = "0.1"
	g.Properties["fontname"] = "Helvetica"
	g.Properties["colorscheme"] = "svg"
	g.Properties["splines"] = "ortho"
	g.Properties["color"] = "#ff00ff"
	g.AddSubGraph(hc)

	fmt.Printf(g.GraphViz())
}
