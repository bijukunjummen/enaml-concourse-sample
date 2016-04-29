package cloudconfig

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml/cloudproperties/aws"
)

const (
	DefaultDiskType         = "gp2"
	DiskSmallName           = "small"
	DiskMediumName          = "medium"
	DiskLargeName           = "large"
	SmallVMName             = "small"
	SmallVMSize             = "t2.micro"
	MediumVMName            = "medium"
	MediumVMSize            = "m3.medium"
	MediumDiskType          = DefaultDiskType
	MediumEphemeralDiskSize = 30000
	SmallDiskType           = DefaultDiskType
	SmallEphemeralDiskSize  = 3000
	AZ1Name                 = "z1"
	AZ2Name                 = "z2"
	PrivateNetworkName      = "private"
	VIPNetworkName          = "vip"
	Region                  = awscloudproperties.USWest
	SubnetPropertyName1     = "subnet-5f423a06"
	SubnetPropertyName2     = "subnet-xxxxxxxx"
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
	cfg.AddDiskType(createDiskType(DiskSmallName, 3000, DefaultDiskType))
	cfg.AddDiskType(createDiskType(DiskMediumName, 20000, DefaultDiskType))
	cfg.AddDiskType(createDiskType(DiskLargeName, 50000, DefaultDiskType))
}

func createDiskType(name string, size int, typename string) enaml.DiskType {
	return enaml.DiskType{
		Name:     name,
		DiskSize: size,
		CloudProperties: awscloudproperties.EphemeralDisk{
			DiskType: typename,
		}}
}

func AddNetwork(cfg *enaml.CloudConfigManifest) {
	octet1 := "10.0.0"
	octet2 := "10.10.0"
	dns := octet1 + ".2"
	privateNetwork := enaml.NewManualNetwork(PrivateNetworkName)
	privateNetwork.AddSubnet(createSubnet(octet1, dns, AZ1Name, SubnetPropertyName1))
	privateNetwork.AddSubnet(createSubnet(octet2, dns, AZ2Name, SubnetPropertyName2))
	cfg.AddNetwork(privateNetwork)
	cfg.AddNetwork(enaml.NewVIPNetwork(VIPNetworkName))
}

func createSubnet(octet, dns, azname, subnetPropertyName string) enaml.Subnet {
	subnet := enaml.NewSubnet(octet, azname)
	subnet.AddDNS(dns)
	subnet.AddReserved(octet + ".1-" + octet + ".10")
	subnet.CloudProperties = awscloudproperties.Network{
		Subnet: subnetPropertyName,
	}
	return subnet
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
