package resources

type Host struct {
	HostName  string
	DNSName   string
	IPAddress string
	Services  []*Service
	State     string
}

func (h *Host) AddService(service *Service) *Host {
	svcInstance := NewService(service.Name, service.Port)
	svcInstance.State = service.State
	h.Services = append(h.Services, svcInstance)

	return h
}

func (h *Host) FindService(name string) *Service {
	for _, s := range h.Services {
		if s.Name == name {
			return s
		}
	}

	return nil
}

func NewHost(hostname string) *Host {
	h := Host{
		HostName: hostname,
		Services: make([]*Service, 0)}

	return &h
}
