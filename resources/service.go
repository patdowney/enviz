package resources

type Service struct {
	Name         string
	Port         int
	Dependencies []*Service
	State        string
}

func (s *Service) AddDependency(service *Service) {
	s.Dependencies = append(s.Dependencies, service)
}

func NewService(name string, port int) *Service {
	return &Service{
		Name:         name,
		Port:         port,
		Dependencies: make([]*Service, 0)}
}
