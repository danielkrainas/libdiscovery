package main

import (
	"fmt"

	"github.com/danielkrainas/libdiscovery"
	_ "github.com/danielkrainas/libdiscovery/locator/consul"
)

func main() {
	d, err := libdiscovery.Create(libdiscovery.CONSUL, map[string]interface{}{})
	if err != nil {
		panic(err)
	}

	nodes, err := d.Nodes()
	if err != nil {
		panic(err)
	}

	for _, n := range nodes {
		fmt.Println("node ", n.Name, "@", n.Address)
	}

	fmt.Printf("node count: %d\n", len(nodes))
}
