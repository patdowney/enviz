package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/patdowney/enviz/json"
	"github.com/patdowney/enviz/resources"
)

func decodeServiceCatalogueFile(filename string) (*resources.ServiceCatalogue, error) {
	catalogueFile, err := os.Open(filename)
	defer catalogueFile.Close()
	if err != nil {
		return nil, err
	}

	catalogue, err := json.DecodeServiceCatalogue(catalogueFile)

	return catalogue, err
}

func decodeHostCatalogueFile(filename string, serviceCatalogue *resources.ServiceCatalogue) (*resources.HostCatalogue, error) {
	catalogueFile, err := os.Open(filename)
	defer catalogueFile.Close()
	if err != nil {
		return nil, err
	}

	catalogue, err := json.DecodeHostCatalogue(catalogueFile, serviceCatalogue)

	return catalogue, err
}

func main() {
	var svcCat string
	flag.StringVar(&svcCat, "service-catalogue", "", "service catalogue")
	flag.Parse()

	svcCatalogue, err := decodeServiceCatalogueFile(svcCat)
	if err != nil {
		fmt.Printf("failed to decode service catalogue: %v\n", err)
		return
	}
	hostCatalogue, err := decodeHostCatalogueFile(flag.Arg(0), svcCatalogue)
	if err != nil {
		fmt.Printf("failed to decode host catalogue: %v\n", err)
		return
	}
	/*
		hosts := make([]*resources.Host, 0)
		for _, arg := range flag.Args() {
			fileHosts := decodeHostsCatalogue(arg)
			for _, h := range fileHosts {
				hosts = append(hosts, h)
			}
		}
	*/
	//	DumpCatalogue(catalogue)
	//graphServices(catalogue)

	graphHosts(hostCatalogue)
}
