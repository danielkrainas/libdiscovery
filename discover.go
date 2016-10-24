package libdiscovery

import (
	"github.com/danielkrainas/libdiscovery/locator"
	"github.com/danielkrainas/libdiscovery/locator/factory"
)

const CONSUL = "consul"

func Create(locatorName string, parameters map[string]interface{}) (*DiscoveryTool, error) {
	driver, err := factory.Create(locatorName, parameters)
	if err != nil {
		return nil, err
	}

	return &DiscoveryTool{driver}, nil
}

type DiscoveryTool struct {
	Driver locator.Driver
}

func (tool *DiscoveryTool) Nodes() ([]*locator.Node, error) {
	return tool.Driver.Nodes()
}

func (tool *DiscoveryTool) Node(name string) (*locator.Node, error) {
	return tool.Driver.Node(name)
}

func (tool *DiscoveryTool) Services() ([]*locator.Service, error) {
	return tool.Driver.Services()
}

func (tool *DiscoveryTool) Service(name string) (*locator.Service, error) {
	return tool.Driver.Service(name)
}
