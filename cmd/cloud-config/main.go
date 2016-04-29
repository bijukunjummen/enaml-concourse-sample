package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/cloud-config"
)

var (
	Version string
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:        "aws",
			Usage:       "aws [--flags]",
			Description: "generate a cloud config manifest for AWS",
			Action: func(c *cli.Context) {
				yamlString, _ := enaml.Cloud(cloudconfig.NewAWSCloudConfig())
				fmt.Println(yamlString)
			},
		},
		{
			Name:        "openstack",
			Usage:       "openstack [--flags]",
			Description: "generate a cloud config manifest for OpenStack",
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:        "vsphere",
			Usage:       "vsphere [--flags]",
			Description: "generate a cloud config manifest for VSphere",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:        "azure",
			Usage:       "azure [--flags]",
			Description: "generate a cloud config manifest for Azure",
			Action: func(c *cli.Context) {
			},
		},
	}
	app.Run(os.Args)
}
