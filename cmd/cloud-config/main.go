package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/enaml-concourse-sample/deployment/cloud-config"
)

var (
	Version string
)

const (
	defaultFileName string = "cloud-config.yml"
)

func main() {
	var err error
	var yamlString string
	app := cli.NewApp()
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:        "aws",
			Usage:       "aws [--flags]",
			Description: "generate a cloud config manifest for AWS",
			Action: func(c *cli.Context) {
				region := c.String("region")
				azList := c.StringSlice("availability-zone")
				subnetList := c.StringSlice("subnet")
				securityGroupList := c.StringSlice("security-group")

				if yamlString, err = enaml.Cloud(cloudconfig.NewAWSCloudConfig(region, azList, subnetList, securityGroupList)); err == nil {
					err = ioutil.WriteFile(defaultFileName, []byte(yamlString), 0644)
					fmt.Println("Generated yaml into", defaultFileName)
				} else {
					panic(err)
				}
			},
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: "subnet", Usage: "subnet name"},
				cli.StringSliceFlag{Name: "availability-zone", Usage: "AZ name"},
				cli.StringFlag{Name: "region", Usage: "aws region"},
				cli.StringSliceFlag{Name: "security-group", Usage: "security group"},
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
