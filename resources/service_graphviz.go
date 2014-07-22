package resources

import (
	"fmt"

	"github.com/patdowney/graphviz"
)

func (s *Service) Node() *graphviz.Node {
	n := graphviz.NewNode(fmt.Sprintf("<f0>%v|<f1>%v", s.Name, s.Port))
	n.ID = fmt.Sprintf("n%p", s)
	n.Properties["shape"] = "Mrecord"

	for _, d := range s.Dependencies {
		n.AddRelation(d.Node())
	}
	// Add Service Details here
	return n
}
