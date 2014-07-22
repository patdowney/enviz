package json

import (
	"encoding/json"
	"io"

	"github.com/patdowney/enviz/resources"
)

type ServiceCatalogue struct {
	Service []resources.Service
}

func TranslateServices(services []*Service) []*resources.Service {
	serviceDefs := make([]*resources.Service, 0)

	for _, svc := range services {
		serviceDefs = append(serviceDefs, resources.NewService(svc.Name, svc.Port))
	}

	return serviceDefs
}

func DecodeServices(inputReader io.Reader) ([]*Service, error) {
	d := json.NewDecoder(inputReader)

	services := make([]*Service, 0)
	err := d.Decode(&services)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func DecodeServiceCatalogue(inputReader io.Reader) (*resources.ServiceCatalogue, error) {

	initialServices, err := DecodeServices(inputReader)
	if err != nil {
		return nil, err
	}

	translatedServices := TranslateServices(initialServices)

	catalogue := resources.NewServiceCatalogue()
	catalogue.AddServices(translatedServices)

	for _, svc := range initialServices {
		err := catalogue.AddServiceDependencies(svc.Name, svc.ServiceRefs)
		if err != nil {
			return nil, err
		}
	}

	return catalogue, nil
}
