package resources

import (
	"fmt"

	"github.com/patdowney/graphviz"
)

func (s *Service) nodeLabel() string {
	label := fmt.Sprintf("<f0>%v", s.Name)
	if s.Port > 0 {
		label = fmt.Sprintf("<f0>%v|<f1>%v", s.Name, s.Port)
	}

	return label
}

func (s *Service) nodeStyle() string {
	line := "solid"
	if s.State == "Pending" {
		line = "dashed"
	}
	return line
}

func (s *Service) Node() *graphviz.Node {

	label := s.nodeLabel()

	n := graphviz.NewNode(label)
	n.ID = fmt.Sprintf("\"n%p\"", s)

	n.Properties["shape"] = "Mrecord"
	n.Properties["style"] = s.nodeStyle()

	for _, d := range s.Dependencies {
		n.AddRelation(d.Node())
	}

	// Add Service Details here
	return n
}
