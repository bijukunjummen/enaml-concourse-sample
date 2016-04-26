package concourse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

var _ = Describe("Concourse Deployment", func() {
	Describe("Given a new deployment", func() {
		XContext("when calling Initialize without a strong password", func() {
			dep := NewDeployment()
			dep.ConcoursePassword = "test"
			It("then we should error and prompt the user for a better pass", func() {
				err := dep.Initialize()
				Ω(err).ShouldNot(BeNil())
			})
		})
	})
})
