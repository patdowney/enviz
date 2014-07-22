package json

import (
	"strings"
	"testing"
)

func TestParseServiceCatalogue(t *testing.T) {
	input := "[ { \"Name\": \"database\", \"Port\": 5432 }, { \"Name\": \"service-a\", \"    Port\": 80, \"ServiceRefs\": [\"database\"] } ]"

	inputReader := strings.NewReader(input)

	services, err := DecodeServices(inputReader)

	if err != nil {
		t.Errorf("failed to decode: %v", err)
		t.FailNow()
	}

	if len(services) != 2 {
		t.Errorf("incorrect length: got:%v, expected:%v", len(services), 2)
		t.FailNow()
	}
}

func TestTranslateServices(t *testing.T) {
	input := []*Service{&Service{
		Name:        "database",
		Port:        1234,
		ServiceRefs: []string{"storage"}}}

	output := TranslateServices(input)

	if len(output[0].Dependencies) != 0 {
		t.Errorf("unexpected Dependencies length: got:%v, expected:%v", output[0].Dependencies, 0)
		t.FailNow()
	}
}

func TestDecodeServiceCatalogue(t *testing.T) {
	input := "[ { \"Name\": \"database\", \"Port\": 5432 }, { \"Name\": \"service-a\", \"    Port\": 80, \"ServiceRefs\": [\"database\"] } ]"

	inputReader := strings.NewReader(input)

	c, err := DecodeServiceCatalogue(inputReader)
	if err != nil {
		t.Errorf("failed to decode catalogue: %v", err)
		t.FailNow()
	}

	if c == nil {
		t.Errorf("failed to create service catalogue")
		t.FailNow()
	}

	svc := c.FindService("service-a")

	if len(svc.Dependencies) != 1 {
		t.Errorf("failed to have correct number of dependencies: got:%v, expected:%v", len(svc.Dependencies), 1)
		t.FailNow()
	}

	ds0 := svc.Dependencies[0]
	if ds0.Port != 5432 {
		t.Errorf("failed to resolve dependency port: got:%v, expected:%v", ds0.Port, 5432)
		t.FailNow()
	}
}
