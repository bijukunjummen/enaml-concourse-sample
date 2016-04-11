package main

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

func main() {
	enaml.Paint(concourse.NewDeployment())
}
