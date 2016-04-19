package resourcepools

import "github.com/xchapter7x/enaml"

var NewFooResource = func(resourceName, networkName string, cpu int) enaml.ResourcePool {

	if cpu > 4 {
		panic("sorry we dont allow you to grab that many cpus")
	}

	return enaml.ResourcePool{
		Name:    resourceName,
		Network: networkName,
		CloudProperties: enaml.CloudProperties{
			"cpu":  cpu,
			"ram":  4096,
			"disk": 10240,
		},
	}
}
