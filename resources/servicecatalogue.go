package resources

import (
	"errors"
	"fmt"
)

type ServiceCatalogue struct {
	serviceIndex map[string]*Service
}

func (c *ServiceCatalogue) AddServiceDependencies(serviceName string, dependencies []string) error {
	parentService := c.FindService(serviceName)

	for _, d := range dependencies {
		ds := c.FindService(d)
		if ds == nil {
			return errors.New(fmt.Sprintf("could not find service dependency: %v", d))
		}

		parentService.AddDependency(ds)
	}
	return nil
}

func (c *ServiceCatalogue) AllServices() []*Service {
	services := make([]*Service, 0)
	for _, svc := range c.serviceIndex {
		services = append(services, svc)
	}
	return services
}

func (c *ServiceCatalogue) AddService(service *Service) {
	c.serviceIndex[service.Name] = service
}

func (c *ServiceCatalogue) AddServices(services []*Service) {
	for _, s := range services {
		c.AddService(s)
	}
}

func (c *ServiceCatalogue) FindService(name string) *Service {
	svc, ok := c.serviceIndex[name]
	if ok {
		return svc
	}

	return nil
}

func NewServiceCatalogue() *ServiceCatalogue {
	return &ServiceCatalogue{serviceIndex: make(map[string]*Service, 0)}
}
