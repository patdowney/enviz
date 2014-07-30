package main

import (
	"fmt"

	"github.com/patdowney/enviz/resources"
	"github.com/patdowney/graphviz"
)

func DumpCatalogue(c *resources.ServiceCatalogue) {
	svcs := c.AllServices()
	for _, s := range svcs {
		fmt.Printf("%v:%d\n", s.Name, s.Port)
		dumpService(s, 1)
		fmt.Printf("\n")
	}
}

func indent(level int, spaces int) {
	for i := 0; i < (level * spaces); i++ {
		fmt.Print(" ")
	}
}

func dumpService(s *resources.Service, level int) {
	for _, d := range s.Dependencies {
		indent(level, 2)
		fmt.Printf("+-> %v:%d\n", d.Name, d.Port)
		dumpService(d, level+1)
	}
}

func graphServices(c *resources.ServiceCatalogue) {
	g := graphviz.NewDigraph("Services")
	g.Properties["ranksep"] = "2"
	g.Properties["fontname"] = "Helvetica"
	//g.Properties["splines"] = "ortho"
	g.AddSubGraph(c)

	fmt.Printf(g.GraphViz())
}
