package cloudconfig

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml/cloudproperties/aws"
)

const (
	SmallVMName            = "small"
	SmallVMSize            = "t2.micro"
	LargeVMName            = "large"
	LargeVMSize            = "m3.medium"
	LargeDiskType          = "gp2"
	LargeEphemeralDiskSize = 30000
	SmallDiskType          = "gp2"
	SmallEphemeralDiskSize = 3000
)

func NewAWSCloudConfig() (awsCloudConfig enaml.CloudConfigManifest) {
	awsCloudConfig = enaml.CloudConfigManifest{}
	AddAZs(&awsCloudConfig)
	AddDisk(&awsCloudConfig)
	AddNetwork(&awsCloudConfig)
	AddVMTypes(&awsCloudConfig)
	return
}

func AddAZs(cfg *enaml.CloudConfigManifest) {
	cfg.AddAZ(enaml.AZ{})
	cfg.AddAZ(enaml.AZ{})
}

func AddDisk(cfg *enaml.CloudConfigManifest) {

}

func AddNetwork(cfg *enaml.CloudConfigManifest) {

}

func AddVMTypes(cfg *enaml.CloudConfigManifest) {
	cfg.AddVMType(enaml.VMType{
		Name: SmallVMName,
		CloudProperties: awscloudproperties.ResourcePool{
			InstanceType: SmallVMSize,
			EphemeralDisk: awscloudproperties.EphemeralDisk{
				Size:     SmallEphemeralDiskSize,
				DiskType: SmallDiskType,
			},
		},
	})
	cfg.AddVMType(enaml.VMType{
		Name: LargeVMName,
		CloudProperties: awscloudproperties.ResourcePool{
			InstanceType: LargeVMSize,
			EphemeralDisk: awscloudproperties.EphemeralDisk{
				Size:     LargeEphemeralDiskSize,
				DiskType: LargeDiskType,
			},
		},
	})
}
