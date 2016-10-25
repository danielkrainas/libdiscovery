package consul

import (
	"github.com/hashicorp/consul/api"

	"github.com/danielkrainas/libdiscovery/locator"
	"github.com/danielkrainas/libdiscovery/locator/factory"
)

func init() {
	factory.Register("consul", &driverFactory{})
}

type driverFactory struct{}

func (factory *driverFactory) Create(parameters map[string]interface{}) (locator.Driver, error) {
	config := api.DefaultConfig()
	if addr, ok := parameters["addr"].(string); ok {
		config.Address = addr
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &driver{client}, nil
}

type driver struct {
	client *api.Client
}

func (d *driver) Nodes() ([]*locator.Node, error) {
	nodes, _, err := d.client.Catalog().Nodes(&api.QueryOptions{})
	if err != nil {
		return nil, err
	}

	results := make([]*locator.Node, len(nodes))
	for i, n := range nodes {
		results[i] = convertNodeData(n)
	}

	return results, nil
}

func (d *driver) Node(name string) (*locator.Node, error) {
	node, _, err := d.client.Catalog().Node(name, &api.QueryOptions{})
	if err != nil {
		return nil, err
	}

	return convertNodeData(node.Node), nil
}

func (d *driver) Services() ([]*locator.Service, error) {
	servicesByTag, _, err := d.client.Catalog().Services(&api.QueryOptions{})
	if err != nil {
		return nil, err
	}

	results := make([]*locator.Service, 0)
	for tag, services := range servicesByTag {
		if tag != "" {
			continue
		}

		for _, name := range services {
			s, err := d.Service(name)
			if err != nil {
				return nil, err
			}

			results = append(results, s)
		}
	}

	return results, nil
}

func (d *driver) Service(name string) (*locator.Service, error) {
	services, _, err := d.client.Catalog().Service(name, "", &api.QueryOptions{})
	if err != nil {
		return nil, err
	}

	if len(services) < 1 {
		return nil, nil
	}

	return convertServiceData(services[0]), nil
}

func convertNodeData(n *api.Node) *locator.Node {
	return &locator.Node{
		Name:    n.Node,
		Address: n.Address,
	}
}

func convertServiceData(s *api.CatalogService) *locator.Service {
	return &locator.Service{
		ID:      s.ServiceID,
		Name:    s.ServiceName,
		Node:    s.Node,
		Port:    uint(s.ServicePort),
		Address: s.Address,
	}
}
