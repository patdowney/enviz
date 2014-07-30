package resources

import (
	"errors"
	"fmt"
)

type HostCatalogue struct {
	hostIndex        map[string]*Host
	ServiceCatalogue *ServiceCatalogue
}

func (c *HostCatalogue) AddHostServices(hostName string, serviceRefs []string) error {
	h := c.FindHost(hostName)
	if h != nil {
		for _, svcName := range serviceRefs {
			svc := c.ServiceCatalogue.FindService(svcName)
			if svc != nil {
				h.AddService(svc)
			} else {
				return errors.New(fmt.Sprintf("No '%v' service found in catalogue", svcName))
			}
		}
	} /* else {
	}*/
	return nil
}

func (c *HostCatalogue) AllHosts() []*Host {
	hosts := make([]*Host, 0)
	for _, h := range c.hostIndex {
		hosts = append(hosts, h)
	}
	return hosts
}

func (c *HostCatalogue) AddHost(host *Host) {
	c.hostIndex[host.HostName] = host
}

func (c *HostCatalogue) AddHosts(hosts []*Host) {
	for _, h := range hosts {
		c.AddHost(h)
	}
}

func (c *HostCatalogue) FindHost(name string) *Host {
	h, ok := c.hostIndex[name]
	if ok {
		return h
	}

	return nil
}

func NewHostCatalogue(serviceCatalogue *ServiceCatalogue) *HostCatalogue {
	return &HostCatalogue{
		hostIndex:        make(map[string]*Host, 0),
		ServiceCatalogue: serviceCatalogue}
}
