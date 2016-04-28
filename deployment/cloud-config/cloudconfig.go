package cloudconfig

import "github.com/xchapter7x/enaml"

func NewCloudConfig(iaas string) (config enaml.CloudConfig) {
	config = NewAWSCloudConfig()
	return
}
