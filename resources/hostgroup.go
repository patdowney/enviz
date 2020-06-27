package resources

type HostGroup struct {
	Name       string
	Hosts      []*Host
	HostGroups []*HostGroup
}

func (g *HostGroup) AddHost(h *Host) *HostGroup {
	g.Hosts = append(g.Hosts, h)
	return g
}

func NewHostGroup(name string) *HostGroup {
	g := HostGroup{
		Name:       name,
		Hosts:      make([]*Host, 0),
		HostGroups: make([]*HostGroup, 0)}

	return &g
}
