package cloudconfig

import "github.com/xchapter7x/enaml"

func NewAWSCloudConfig() (awsCloudConfig enaml.CloudConfigManifest) {
	awsCloudConfig = enaml.CloudConfigManifest{}
	AddAZs(&awsCloudConfig)
	return
}

func AddAZs(cfg *enaml.CloudConfigManifest) {
	cfg.AZs = append(cfg.AZs, enaml.AZ{})
	cfg.AZs = append(cfg.AZs, enaml.AZ{})
}
