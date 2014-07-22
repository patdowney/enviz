package resources

type HostGroup struct {
	Name  string
	Hosts []*Host
}

func (g *HostGroup) AddHost(h *Host) *HostGroup {
	g.Hosts = append(g.Hosts, h)
	return g
}

func NewHostGroup(name string) *HostGroup {
	g := HostGroup{
		Name:  name,
		Hosts: make([]*Host, 0)}

	return &g
}

type Host struct {
	HostName  string
	DNSName   string
	IPAddress string
	Services  []*Service
}

func (h *Host) AddService(service *Service) *Host {
	h.Services = append(h.Services, service)

	return h
}

func NewHost(hostname string) *Host {
	h := Host{
		HostName: hostname,
		Services: make([]*Service, 0)}

	return &h
}
