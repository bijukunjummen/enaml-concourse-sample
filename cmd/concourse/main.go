package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

var (
	//Version -
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "Concourse Yaml"
	app.Usage = "cli generator"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			//Usage:       "--output-file-name <filename> --adminuser <usr> --adminpass <pass> --opsmanageruser <opsuser> --opsmanagerpass <opspass> -d <dir> --tile elastic-runtime",
			Description: "generate concourse yaml for given arguments",
			Action:      generate,
			Flags:       generateFlags(),
		},
	}
	app.Run(os.Args)
}

const (
	defaultFileName          string = "concourse.yml"
	outputFileName           string = "OUTPUT_FILENAME"
	concoursePassword        string = "PASSWORD"
	concourseUsername        string = "USERNAME"
	concourseURL             string = "URL"
	concourseWebInstances    string = "WEB_INSTANCES"
	concourseWebIPs          string = "WEB_IPS"
	boshDirectorUUID         string = "BOSH_DIRECTOR_UUID"
	boshStemcellAlias        string = "BOSH_STEMCELL_ALIAS"
	concourseNetworkName     string = "NETWORK_NAME"
	concourseNetworkRange    string = "NETWORK_RANGE"
	concourseNetworkGateway  string = "NETWORK_GATEWAY"
	concourseWebAZs          string = "WEB_AZS"
	concourseDatabaseAZs     string = "DATABASE_AZS"
	concourseWorkerAZs       string = "WORKER_AZS"
	boshCloudConfig          string = "BOSH_CLOUD_CONFIG"
	concourseDeploymentName  string = "BOSH_DEPLOYMENT_NAME"
	concoursePostgresqlDbPwd string = "POSTGRESQL_DB_PWD"
)

func getFlag(input string) (flag string) {
	flag = strings.ToLower(strings.Replace(input, "_", "-", -1))
	return
}
func generateFlags() (flags []cli.Flag) {
	var flagList = map[string]flagBucket{
		outputFileName: flagBucket{
			Desc:   "destination for output",
			EnvVar: outputFileName,
		},
		boshCloudConfig: flagBucket{
			Desc:   "true/false for generate cloudConfig compatible (default true)",
			EnvVar: boshCloudConfig,
		},
		boshDirectorUUID: flagBucket{
			Desc:   "bosh director uuid (bosh status --uuid)",
			EnvVar: boshDirectorUUID,
		},
		boshStemcellAlias: flagBucket{
			Desc:   "url for concourse ui",
			EnvVar: boshStemcellAlias,
		},
		concourseWebInstances: flagBucket{
			Desc:   "number of web instances (default 1)",
			EnvVar: concourseWebInstances,
		},
		concourseWebIPs: flagBucket{
			Desc:        "array of IPs reserved for concourse web-ui",
			EnvVar:      concourseWebIPs,
			StringSlice: true,
		},
		concourseURL: flagBucket{
			Desc:   "url for concourse ui",
			EnvVar: concourseURL,
		},
		concourseUsername: flagBucket{
			Desc:   "username for concourse ui",
			EnvVar: concourseUsername,
		},
		concoursePassword: flagBucket{
			Desc:   "password for concourse ui",
			EnvVar: concoursePassword,
		},
		concourseNetworkName: flagBucket{
			Desc:   "name of network to deploy concourse on",
			EnvVar: concourseNetworkName,
		},
		concourseNetworkRange: flagBucket{
			Desc:   "network range to deploy concourse on - only applies in non-cloud config",
			EnvVar: concourseNetworkRange,
		},
		concourseNetworkGateway: flagBucket{
			Desc:   "network gateway for concourse - only applies in non-cloud config",
			EnvVar: concourseNetworkGateway,
		},
		concourseWebAZs: flagBucket{
			Desc:        "array of AZs to deploy concourse web jobs to",
			EnvVar:      concourseWebAZs,
			StringSlice: true,
		},
		concourseDatabaseAZs: flagBucket{
			Desc:        "array of AZs to deploy concourse database jobs to",
			EnvVar:      concourseDatabaseAZs,
			StringSlice: true,
		},
		concourseWorkerAZs: flagBucket{
			Desc:        "array of AZs to deploy concourse worker jobs to",
			EnvVar:      concourseWorkerAZs,
			StringSlice: true,
		},
		concourseDeploymentName: flagBucket{
			Desc:   "bosh deployment name",
			EnvVar: concourseDeploymentName,
		},
		concoursePostgresqlDbPwd: flagBucket{
			Desc:   "password for postgresql job",
			EnvVar: concoursePostgresqlDbPwd,
		},
	}
	for _, v := range flagList {
		if v.StringSlice {
			flags = append(flags, cli.StringSliceFlag{
				Name:   getFlag(v.EnvVar),
				Usage:  v.Desc,
				EnvVar: v.EnvVar,
			})
		} else {
			flags = append(flags, cli.StringFlag{
				Name:   getFlag(v.EnvVar),
				Value:  "",
				Usage:  v.Desc,
				EnvVar: v.EnvVar,
			})
		}
	}
	return
}

type flagBucket struct {
	Desc        string
	EnvVar      string
	StringSlice bool
}

func generate(c *cli.Context) {
	var fileName string

	fileName = c.String(getFlag(outputFileName))
	if fileName == "" {
		fileName = defaultFileName
	}
	deployment := concourse.NewDeployment()
	if c.IsSet(getFlag(concourseDeploymentName)) {
		deployment.DeploymentName = c.String(getFlag(concourseDeploymentName))
	} else {
		deployment.DeploymentName = "concourse"
	}

	if c.IsSet(getFlag(concoursePostgresqlDbPwd)) {
		deployment.PostgresPassword = c.String(getFlag(concoursePostgresqlDbPwd))
	} else {
		deployment.PostgresPassword = "dummy-postgres-password"
	}
	deployment.ConcoursePassword = c.String(getFlag(concoursePassword))
	deployment.ConcourseUserName = c.String(getFlag(concourseUsername))
	deployment.ConcourseURL = c.String(getFlag(concourseURL))
	deployment.DirectorUUID = c.String(getFlag(boshDirectorUUID))
	deployment.StemcellAlias = c.String(getFlag(boshStemcellAlias))
	deployment.NetworkName = c.String(getFlag(concourseNetworkName))
	if c.IsSet(getFlag(boshCloudConfig)) {
		deployment.CloudConfig = c.Bool(getFlag(boshCloudConfig))
	} else {
		deployment.CloudConfig = true
	}
	if c.IsSet(getFlag(concourseWebInstances)) {
		deployment.WebInstances = c.Int(getFlag(concourseWebInstances))
	} else {
		deployment.WebInstances = 1

	}
	if c.IsSet(getFlag(concourseWebIPs)) {
		deployment.WebIPs = c.StringSlice(getFlag(concourseWebIPs))
	}
	deployment.NetworkRange = c.String(getFlag(concourseNetworkRange))
	deployment.NetworkGateway = c.String(getFlag(concourseNetworkGateway))

	if c.IsSet(getFlag(concourseWebAZs)) {
		deployment.WebAZs = c.StringSlice(getFlag(concourseWebAZs))
	}
	if c.IsSet(getFlag(concourseDatabaseAZs)) {
		deployment.DatabaseAZs = c.StringSlice(getFlag(concourseDatabaseAZs))
	}
	if c.IsSet(getFlag(concourseWorkerAZs)) {
		deployment.WorkerAZs = c.StringSlice(getFlag(concourseWorkerAZs))
	}

	var yamlString string
	var err error
	if err = deployment.Initialize(); err == nil {
		if yamlString, err = enaml.Paint(deployment); err == nil {
			err = ioutil.WriteFile(fileName, []byte(yamlString), 0644)
		}
	}
	if err != nil {
		panic(err.Error())
	}
	println("completed generating yaml into ", fileName)
}
