# libdiscovery

[![License](https://img.shields.io/badge/license-Unlicense-blue.svg?style=flat)](UNLICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/danielkrainas/libdiscovery)](https://goreportcard.com/report/github.com/danielkrainas/libdiscovery)

## Installation

> $ go get github.com/danielkrainas/libdiscovery

## Usage

Code usage example:

```
package main

import (
	"fmt"

	"github.com/danielkrainas/libdiscovery"

	// include the backend you wish to have available
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
```

## Bugs and Feedback

If you see a bug or have a suggestion, feel free to open an issue [here](https://github.com/danielkrainas/libdiscovery/issues).

## Contributions

PR's welcome! There are no strict style guidelines, just follow best practices and try to keep with the general look & feel of the code present. All submissions should atleast be `go fmt -s` and have a test to verify *(if applicable)*.

For details on how to extend and develop libdiscovery, see the [development guide](docs/development).

## License

[Unlicense](http://unlicense.org/UNLICENSE). This is a Public Domain work. 

[![Public Domain](https://licensebuttons.net/p/mark/1.0/88x31.png)](http://questioncopyright.org/promise)

> ["Make art not law"](http://questioncopyright.org/make_art_not_law_interview) -Nina Paley