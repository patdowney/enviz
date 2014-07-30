package json

type Service struct {
	Name           string
	Port           int
	DependencyRefs []string
	State          string
}
