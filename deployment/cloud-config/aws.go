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
	AZ1Name                = "az1"
	AZ2Name                = "az2"
)

func NewAWSCloudConfig() (awsCloudConfig *enaml.CloudConfigManifest) {
	awsCloudConfig = &enaml.CloudConfigManifest{}
	AddAZs(awsCloudConfig)
	AddDisk(awsCloudConfig)
	AddNetwork(awsCloudConfig)
	AddVMTypes(awsCloudConfig)
	return
}

func AddAZs(cfg *enaml.CloudConfigManifest) {
	cfg.AddAZ(enaml.AZ{
		Name: AZ1Name,
		CloudProperties: awscloudproperties.AZ{
			AvailabilityZoneName: awscloudproperties.USEast + "a",
		},
	})
	cfg.AddAZ(enaml.AZ{
		Name: AZ2Name,
		CloudProperties: awscloudproperties.AZ{
			AvailabilityZoneName: awscloudproperties.USEast + "b",
		},
	})
}

func AddDisk(cfg *enaml.CloudConfigManifest) {

}

func AddNetwork(cfg *enaml.CloudConfigManifest) {

}

func AddVMTypes(cfg *enaml.CloudConfigManifest) {
	cfg.AddVMType(enaml.VMType{
		Name:            SmallVMName,
		CloudProperties: NewVMCloudProperty(SmallVMSize, SmallDiskType, SmallEphemeralDiskSize),
	})
	cfg.AddVMType(enaml.VMType{
		Name:            LargeVMName,
		CloudProperties: NewVMCloudProperty(LargeVMSize, LargeDiskType, LargeEphemeralDiskSize),
	})
}

func NewVMCloudProperty(instanceType, diskType string, diskSize int) awscloudproperties.VMType {
	return awscloudproperties.VMType{
		InstanceType: instanceType,
		EphemeralDisk: awscloudproperties.EphemeralDisk{
			Size:     diskSize,
			DiskType: diskType,
		},
	}
}
