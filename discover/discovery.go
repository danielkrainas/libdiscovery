package libdiscovery

import ()

type Service struct {
	Name    string `json:"name"`
	Address string `json:"name"`
	Port    uint   `json:"port"`
}
