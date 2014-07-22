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

	s := NewClusterSubGraph("SIT")

	n2 := NewNode("node_two")

	n1 := NewNode("node_one")
	r := n1.AddRelation(n2)
	s.AddNode(n1)
	s.AddNode(n2)

	r.Properties["color"] = "blue"

	s1 := NewClusterSubGraph("Zone 3")
	s1.Properties["rankdir"] = "LR"

	n4 := NewNode("node_four")
	n3 := NewNode("node_three")
	r1 := n3.AddRelation(n4)
	r1.Properties["style"] = "invis"

	s1.AddNode(n3)
	s1.AddNode(n4)
	s.AddSubGraph(s1)
	g.AddSubGraph(s)

	log.Print(g.GraphViz())
}
