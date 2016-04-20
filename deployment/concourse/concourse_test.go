package concourse_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

var _ = Describe("Concourse Deployment", func() {
	Describe("Given a new deployment", func() {
		Context("When created using NewDeployment() and a non HA config", func() {
			var dep Deployment
			BeforeEach(func() {
				dep = NewDeployment(1, "", "", "very very long password")
			})
			It("Then the web group should be made always be highly available", func() {
				for i, _ := range dep.Manifest.InstanceGroups {
					if dep.Manifest.InstanceGroups[i].Name == "web" {
						fmt.Println(dep.Manifest.InstanceGroups[i].Instances)
						Ω(dep.Manifest.InstanceGroups[i].Instances).Should(BeNumerically(">", 1))
					}
				}
			})
		})

		Context("when calling NewDeployment without a strong password", func() {
			It("then we should panic and prompt the user for a better pass", func() {
				Ω(func() {
					NewDeployment(3, "", "", "")
				}).Should(Panic())
			})
		})
	})
})
