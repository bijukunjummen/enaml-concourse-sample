package cloudconfig

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml/cloudproperties/aws"
)

const (
	SmallVMName             = "small"
	SmallVMSize             = "t2.micro"
	MediumVMName            = "medium"
	MediumVMSize            = "m3.medium"
	MediumDiskType          = "gp2"
	MediumEphemeralDiskSize = 30000
	SmallDiskType           = "gp2"
	SmallEphemeralDiskSize  = 3000
	AZ1Name                 = "z1"
	AZ2Name                 = "z2"
	PrivateNetworkName      = "private"
	VIPNetworkName          = "vip"
	Region                  = awscloudproperties.USWest
)

var (
	DefaultSecurityGroups = []string{"bosh-cwashburn-InternalSecurityGroup-WQAFGW1Z5W0Y"}
)

func NewAWSCloudConfig() (awsCloudConfig *enaml.CloudConfigManifest) {
	awsCloudConfig = &enaml.CloudConfigManifest{}
	AddAZs(awsCloudConfig)
	AddDisk(awsCloudConfig)
	AddNetwork(awsCloudConfig)
	AddVMTypes(awsCloudConfig)
	AddCompilation(awsCloudConfig, AZ1Name, MediumVMName, PrivateNetworkName)
	return
}

func AddCompilation(cfg *enaml.CloudConfigManifest, az string, vmtype string, network string) {
	cfg.SetCompilation(&enaml.Compilation{
		Workers:             5,
		ReuseCompilationVMs: true,
		AZ:                  az,
		VMType:              vmtype,
		Network:             network,
	})
}

func AddAZs(cfg *enaml.CloudConfigManifest) {
	cfg.AddAZ(enaml.AZ{
		Name: AZ1Name,
		CloudProperties: awscloudproperties.AZ{
			AvailabilityZoneName: Region + "a",
			SecurityGroups:       DefaultSecurityGroups,
		},
	})
	cfg.AddAZ(enaml.AZ{
		Name: AZ2Name,
		CloudProperties: awscloudproperties.AZ{
			AvailabilityZoneName: Region + "b",
			SecurityGroups:       DefaultSecurityGroups,
		},
	})
}

func AddDisk(cfg *enaml.CloudConfigManifest) {

}

func AddNetwork(cfg *enaml.CloudConfigManifest) {
	octet := "10.0.0"
	dns := octet + ".2"
	privateNetwork := enaml.NewManualNetwork(PrivateNetworkName)
	privateSubnet := enaml.NewSubnet(octet, AZ1Name)
	privateSubnet.AddDNS(dns)
	privateNetwork.AddSubnet(privateSubnet)
	privateNetwork.AddSubnet(enaml.Subnet{})
	cfg.AddNetwork(privateNetwork)
	cfg.AddNetwork(enaml.NewVIPNetwork(VIPNetworkName))
}

func AddVMTypes(cfg *enaml.CloudConfigManifest) {
	cfg.AddVMType(enaml.VMType{
		Name:            SmallVMName,
		CloudProperties: NewVMCloudProperty(SmallVMSize, SmallDiskType, SmallEphemeralDiskSize),
	})
	cfg.AddVMType(enaml.VMType{
		Name:            MediumVMName,
		CloudProperties: NewVMCloudProperty(MediumVMSize, MediumDiskType, MediumEphemeralDiskSize),
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
