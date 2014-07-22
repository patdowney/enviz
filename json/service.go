package json

type Service struct {
	Name        string
	Port        int
	ServiceRefs []string
	//DependencyRefs []string
}
