package main

import (
	"fmt"
	"os"
	"strings"

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
				region := c.String("region")
				azList := strings.Split(c.String("availability-zones"), ",")
				subnetList := strings.Split(c.String("subnets"), ",")
				securityGroupList := strings.Split(c.String("security-groups"), ",")

				if yamlString, err := enaml.Cloud(cloudconfig.NewAWSCloudConfig(region, azList, subnetList, securityGroupList)); err == nil {
					fmt.Println(yamlString)
				} else {
					fmt.Println(err)
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "subnets", Usage: "comma separated list of subnet names"},
				cli.StringFlag{Name: "availability-zones", Usage: "comma separated list of AZ names"},
				cli.StringFlag{Name: "region", Usage: "aws region"},
				cli.StringFlag{Name: "security-groups", Usage: "comma separated list of security groups"},
			},
		},
		{
			Name:        "openstack",
			Usage:       "openstack [--flags]",
			Description: "generate a cloud config manifest for OpenStack",
			Action: func(c *cli.Context) {
				fmt.Println("openstack iaas is not implemented yet")
			},
		},
		{
			Name:        "vsphere",
			Usage:       "vsphere [--flags]",
			Description: "generate a cloud config manifest for VSphere",
			Action: func(c *cli.Context) {
				fmt.Println("vsphere iaas is not implemented yet")
			},
		},
		{
			Name:        "azure",
			Usage:       "azure [--flags]",
			Description: "generate a cloud config manifest for Azure",
			Action: func(c *cli.Context) {
				fmt.Println("azure iaas is not implemented yet")
			},
		},
	}
	app.Run(os.Args)
}
