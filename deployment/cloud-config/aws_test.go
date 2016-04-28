package cloudconfig_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/enaml"
	. "github.com/xchapter7x/enaml-concourse-sample/deployment/cloud-config"
	"github.com/xchapter7x/enaml/cloudproperties/aws"
)

var _ = Describe("given CloudConfig Deployment for AWS", func() {
	var awsConfig enaml.CloudConfigManifest
	BeforeEach(func() {
		awsConfig = NewAWSCloudConfig()
	})

	Context("when AZs are defined", func() {
		It("then we should always have a HA configuration", func() {
			haCount := 2
			numberOfAZs := len(awsConfig.AZs)
			Ω(numberOfAZs).Should(BeNumerically(">=", haCount))
		})
	})
	Context("when a user of the iaas would like to define vm types", func() {
		It("then there should be 2 vm type options available", func() {
			Ω(len(awsConfig.VMTypes)).Should(Equal(2))
		})

		It("then they should have the option of a small VM configuration", func() {
			_, err := getVmTypeByName(SmallVMName, awsConfig.VMTypes)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("when the vmtype is small", func() {
			var vm enaml.VMType
			BeforeEach(func() {
				vm, _ = getVmTypeByName(SmallVMName, awsConfig.VMTypes)
			})
			It("then it should use a t2.micro size aws instance", func() {
				Ω(vm.CloudProperties.(awscloudproperties.VMType).InstanceType).Should(Equal(SmallVMSize))
			})

			It("then it should use a properly configured ephemeral disk", func() {
				properSmallDiskSize := SmallEphemeralDiskSize
				properDiskType := SmallDiskType
				Ω(vm.CloudProperties.(awscloudproperties.VMType).EphemeralDisk.Size).Should(Equal(properSmallDiskSize))
				Ω(vm.CloudProperties.(awscloudproperties.VMType).EphemeralDisk.DiskType).Should(Equal(properDiskType))
			})
		})

		It("then they should have the option of a large VM configuration", func() {
			_, err := getVmTypeByName(LargeVMName, awsConfig.VMTypes)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("when the vmtype is large", func() {
			var vm enaml.VMType
			BeforeEach(func() {
				vm, _ = getVmTypeByName(LargeVMName, awsConfig.VMTypes)
			})
			It("then it should use a m3.medium size aws instance", func() {
				Ω(vm.CloudProperties.(awscloudproperties.VMType).InstanceType).Should(Equal(LargeVMSize))
			})

			It("then it should use a properly configured ephemeral disk", func() {
				properLargeDiskSize := LargeEphemeralDiskSize
				properDiskType := LargeDiskType
				Ω(vm.CloudProperties.(awscloudproperties.VMType).EphemeralDisk.Size).Should(Equal(properLargeDiskSize))
				Ω(vm.CloudProperties.(awscloudproperties.VMType).EphemeralDisk.DiskType).Should(Equal(properDiskType))
			})
		})
	})

	XContext("when a user of the iaas would like to assign disk", func() {
		It("then they should have the option of small or large capacity configurations", func() {

			Ω(1).Should(Equal(1))
		})
	})

	XContext("when a user of the iaas would like to assign a network", func() {
		It("then they should have the option of a private and a vip network", func() {

			Ω(1).Should(Equal(1))
		})

		Context("when a user of the iaas is assigning the private network", func() {

			It("then they should have the option of multiple subnets - one for each configured AZ", func() {

				Ω(1).Should(Equal(1))
			})
		})
	})
})

func getVmTypeByName(name string, vmTypes []enaml.VMType) (res enaml.VMType, err error) {
	err = errors.New("no type found")
	for _, k := range vmTypes {
		if k.Name == name {
			err = nil
			res = k
		}
	}
	return
}
