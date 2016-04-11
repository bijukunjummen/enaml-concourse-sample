package resourcepools

import "github.com/xchapter7x/enaml"

var NewFooResource = func(resourceName, networkName string) enaml.ResourcePool {

	return enaml.ResourcePool{
		Name:    resourceName,
		Network: networkName,
		CloudProperties: enaml.CloudProperties{
			"cpu":  2,
			"ram":  4096,
			"disk": 10240,
		},
	}
}
