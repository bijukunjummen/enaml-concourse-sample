package main

import (
	"fmt"
	"os"

	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/cloud-config"
)

func main() {
	var iaas = os.Getenv("CLOUDCONFIG_IAAS")

	yamlString, _ := enaml.Cloud(cloudconfig.NewCloudConfig(iaas))
	fmt.Println(yamlString)
}
