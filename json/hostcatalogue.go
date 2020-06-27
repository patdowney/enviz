package json

import (
	"encoding/json"
	"io"

	"github.com/patdowney/enviz/resources"
)

type HostCatalogue struct {
	Host []resources.Host
}

func TranslateHosts(hosts []*Host) []*resources.Host {
	hostDefs := make([]*resources.Host, 0)

	for _, host := range hosts {
		h := resources.NewHost(host.HostName)
		h.DNSName = host.DNSName
		h.IPAddress = host.IPAddress
		h.State = host.State

		hostDefs = append(hostDefs, h)
	}

	return hostDefs
}

func DecodeHosts(inputReader io.Reader) ([]*Host, error) {
	d := json.NewDecoder(inputReader)

	hosts := make([]*Host, 0)
	err := d.Decode(&hosts)

	if err != nil {
		return nil, err
	}

	return hosts, nil
}

func DecodeHostCatalogue(inputReader io.Reader, serviceCatalogue *resources.ServiceCatalogue) (*resources.HostCatalogue, error) {

	initialHosts, err := DecodeHosts(inputReader)
	if err != nil {
		return nil, err
	}

	translatedHosts := TranslateHosts(initialHosts)

	catalogue := resources.NewHostCatalogue(serviceCatalogue)
	catalogue.AddHosts(translatedHosts)

	for _, h := range initialHosts {
		err := catalogue.AddHostServices(h.HostName, h.ServiceRefs)
		if err != nil {
			return nil, err

		}

		err = catalogue.AddUpstreamHosts(h.HostName, h.PrimaryUpstreamHosts, h.SecondaryUpstreamHosts)
		if err != nil {
			return nil, err

		}
	}

	return catalogue, nil
}
