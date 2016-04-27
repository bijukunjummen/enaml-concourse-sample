package concourse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

var _ = Describe("Concourse Deployment", func() {
	var deployment Deployment
	BeforeEach(func() {
		deployment = NewDeployment()
	})
	Describe("Given a CreateCompilation", func() {
		Context("when calling with a networkName", func() {
			It("then we should return a valid enamel.Compilation", func() {
				networkName := "private"
				compilation := deployment.CreateCompilation(networkName)
				Ω(compilation.Network).Should(Equal(networkName))
				Ω(compilation.Workers).Should(Equal(3))
				Ω(compilation.ReuseCompilationVMs).Should(Equal(false))
			})
		})
	})
	Describe("Given a CreateResourcePool", func() {
		Context("when calling with a networkName", func() {
			It("then we should return a valid enamel.ResourcePool", func() {
				networkName := "private"
				resourcePool := deployment.CreateResourcePool(networkName)
				Ω(resourcePool.Network).Should(Equal(networkName))
				Ω(resourcePool.Name).Should(Equal("concourse"))
				stemcell := resourcePool.Stemcell
				Ω(stemcell.Name).Should(Equal("bosh-warden-boshlite-ubuntu-trusty-go_agent"))
				Ω(stemcell.Version).Should(Equal("latest"))
				Ω(stemcell.OS).Should(Equal(""))
				Ω(stemcell.SHA1).Should(Equal(""))
				Ω(stemcell.URL).Should(Equal(""))
				Ω(stemcell.Alias).Should(Equal(""))
			})
		})
	})
	Describe("Given a CreateUpdate", func() {
		Context("when calling", func() {
			It("then we should return a valid enamel.Update", func() {
				update := deployment.CreateUpdate()
				Ω(update.Canaries).Should(Equal(1))
				Ω(update.MaxInFlight).Should(Equal(3))
				Ω(update.Serial).Should(Equal(false))
				Ω(update.UpdateWatchTime).Should(Equal("1000-60000"))
				Ω(update.CanaryWatchTime).Should(Equal("1000-60000"))
			})
		})
	})

	Describe("Given a CreateManualDeploymentNetwork", func() {
		Context("when calling with network name", func() {
			It("then we should return a valid enamel.ManualNetwork", func() {
				networkName := "private"
				networkRange := "10.0.0.0/24"
				networkGateway := "10.244.8.1"
				webIPs := []string{"10.244.8.2"}
				manualNetwork := deployment.CreateManualDeploymentNetwork(networkName, networkRange, networkGateway, webIPs)
				Ω(manualNetwork.Name).Should(Equal(networkName))
				Ω(manualNetwork.Type).Should(Equal("manual"))
				subnets := manualNetwork.Subnets
				Ω(len(subnets)).Should(Equal(1))
				Ω(subnets[0].Range).Should(Equal(networkRange))
				Ω(subnets[0].Gateway).Should(Equal(networkGateway))
				Ω(subnets[0].Static).Should(Equal(webIPs))
				Ω(subnets[0].DNS).Should(Equal(""))
				Ω(subnets[0].Reserved).Should(BeEmpty())
				Ω(subnets[0].AZs).Should(BeEmpty())
				Ω(subnets[0].AZ).Should(Equal(""))
			})
		})
	})
	Describe("Given a new deployment", func() {
		XContext("when calling Initialize without a strong password", func() {

			deployment.ConcoursePassword = "test"
			It("then we should error and prompt the user for a better pass", func() {
				err := deployment.Initialize()
				Ω(err).ShouldNot(BeNil())
			})
		})
	})
})
