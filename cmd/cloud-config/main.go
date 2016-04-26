package main

import (
	"os"

	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/cloud-config"
)

func main() {
	var iaas = os.Getenv("CLOUDCONFIG_IAAS")
	enaml.Paint(cloudconfig.NewDeployment(iaas))
}
