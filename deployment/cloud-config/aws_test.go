package cloudconfig_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/enaml"
	. "github.com/xchapter7x/enaml-concourse-sample/deployment/cloud-config"
)

var _ = XDescribe("given CloudConfig Deployment for AWS", func() {
	var awsConfig enaml.DeploymentManifest
	BeforeEach(func() {
		awsConfig = NewAWSCloudConfig()
		fmt.Println(awsConfig)
	})

	Context("when AZs are defined", func() {
		It("then we should always have a HA configuration", func() {
			Ω(1).Should(Equal(1))
		})
	})
	Context("when a user of the iaas would like to define vm types", func() {
		It("then they should have the option of a small and a large VM configuration", func() {
			Ω(1).Should(Equal(1))
		})
	})
	Context("when a user of the iaas would like to assign disk", func() {
		It("then they should have the option of small or large capacity configurations", func() {

			Ω(1).Should(Equal(1))
		})
	})

	Context("when a user of the iaas would like to assign a network", func() {
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
