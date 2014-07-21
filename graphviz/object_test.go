package graphviz

import (
	"log"
	"testing"
)

func TestTest(t *testing.T) {

	g := NewDigraph("G")
	g.Properties["fontname"] = "Helvetica"
	a := NewAttr("node")
	a.Properties["label"] = "somelabel"
	a.Properties["width"] = "2"
	a.Properties["shape"] = "Mrecord"

	g.AddAttribute(a)

	s := NewSubGraph("clusterBlah")
	s.Properties["label"] = "SIT"

	s.AddNode(NewNode("node_one"))
	s.Nodes[0].Properties["label"] = "node_one"

	s.AddNode(NewNode("node_two"))
	s.Nodes[1].Properties["label"] = "node_two"

	s1 := NewSubGraph("clusterBlah2")
	s1.Properties["label"] = "Zone 3"
	s1.AddNode(&Node{Name: "node_three"})
	s1.AddNode(&Node{Name: "node_four"})
	s.AddSubGraph(s1)
	g.AddSubGraph(s)

	r := NewRelation("node_one", "node_two")
	r.Properties["color"] = "blue"
	g.AddRelation(r)

	log.Print(g.GraphViz())
}
