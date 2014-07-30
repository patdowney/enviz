package resources

import "github.com/patdowney/graphviz"

func (c *HostCatalogue) FindHostsWithService(name string) []*Host {
	hosts := make([]*Host, 0)

	for _, h := range c.AllHosts() {
		for _, svc := range h.Services {
			if svc.Name == name {
				hosts = append(hosts, h)
			}
		}
	}

	return hosts
}

type ServiceConnection struct {
	SourceHost    *Host
	SourceService *Service
	TargetHost    *Host
	TargetService *Service
}

func (c *ServiceConnection) edgeStyle() string {
	style := "solid"
	if c.SourceService.State == "Pending" || c.TargetService.State == "Pending" {
		style = "dotted"
	}
	return style
}

func (c *ServiceConnection) Relation() *graphviz.Relation {
	sourceNode := c.SourceService.Node()
	targetNode := c.TargetService.Node()
	r := graphviz.NewRelation(
		sourceNode,
		targetNode)

	r.Properties["style"] = c.edgeStyle()

	return r
}

func (c *HostCatalogue) FindServiceConnections(host *Host) []*ServiceConnection {
	connections := make([]*ServiceConnection, 0)

	//host := c.hostIndex[h.HostName]

	//  for each hosts services
	for _, hostService := range host.Services {
		//    lookup service in service catalogue,
		hostServiceDependencies := c.ServiceCatalogue.FindService(hostService.Name).Dependencies
		//    for each service dependency
		for _, serviceDep := range hostServiceDependencies {
			//    find hosts with services
			serviceHosts := c.FindHostsWithService(serviceDep.Name)

			for _, targetHost := range serviceHosts {
				targetService := targetHost.FindService(serviceDep.Name)
				c := &ServiceConnection{
					SourceHost:    host,
					SourceService: hostService,
					TargetHost:    targetHost,
					TargetService: targetService}
				connections = append(connections, c)
			}
		}
	}

	return connections
}

func (c *HostCatalogue) Graph() *graphviz.GraphBase {
	g := graphviz.NewClusterSubGraph("Hosts")

	for _, h := range c.hostIndex {
		g.AddSubGraph(h)
	}

	/* find relations */
	// for each host
	//	for _, h := range c.AllHosts() {
	for _, h := range c.hostIndex {
		hostConns := c.FindServiceConnections(h)
		for _, conn := range hostConns {
			g.AddRelation(conn)
		}
	}

	return g
}
