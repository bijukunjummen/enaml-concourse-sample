package networks

import "github.com/xchapter7x/enaml"

var NewFooNetworkExternal = func(networkName string) enaml.ManualNetwork {
	return enaml.ManualNetwork{
		Name: networkName,
		Type: "manual",
		Subnets: []enaml.Subnet{
			enaml.Subnet{
				Range:   "10.0.0.0/24",
				DNS:     "10.0.0.2",
				Gateway: "10.0.0.1",
				CloudProperties: enaml.CloudProperties{
					"name": "NETWORK_NAME",
				},
			},
		},
	}
}
