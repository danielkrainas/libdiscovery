package locator

import ()

type Node struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Service struct {
	Name    string `json:"name"`
	Node    string `json:"name"`
	Address string `json:"address"`
	Port    uint   `json:"port"`
}

type Driver interface {
	Nodes() ([]*Node, error)
	Node(name string) (*Node, error)
	Services() ([]*Service, error)
	Service(serviceName string) (*Service, error)
}
