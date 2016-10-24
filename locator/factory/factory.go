package factory

import (
	"fmt"

	"github.com/danielkrainas/libdiscovery/locator"
)

var driverFactories = make(map[string]LocatorDriverFactory)

type LocatorDriverFactory interface {
	Create(parameters map[string]interface{}) (locator.Driver, error)
}

func Register(name string, factory LocatorDriverFactory) {
	if factory == nil {
		panic("LocatorDriverFactory cannot be nil")
	}

	if _, registered := driverFactories[name]; registered {
		panic(fmt.Sprintf("LocatorDriverFactory named %s already registered", name))
	}

	driverFactories[name] = factory
}

func Create(name string, parameters map[string]interface{}) (locator.Driver, error) {
	if factory, ok := driverFactories[name]; ok {
		return factory.Create(parameters)
	}

	return nil, InvalidLocatorDriverError{name}
}

type InvalidLocatorDriverError struct {
	Name string
}

func (err InvalidLocatorDriverError) Error() string {
	return fmt.Sprintf("Locator driver not registered: %s", err.Name)
}
