package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

func main() {
	deployment := concourse.NewDeployment()
	deployment.ConcoursePassword = os.Getenv("CONCOURSE_PASSWORD")
	deployment.ConcourseUserName = os.Getenv("CONCOURSE_USERNAME")
	deployment.ConcourseURL = os.Getenv("CONCOURSE_URL")
	deployment.DirectorUUID = os.Getenv("BOSH_DIRECTOR_UUID")
	deployment.StemcellAlias = os.Getenv("BOSH_STEMCELL_ALIAS")
	deployment.NetworkName = os.Getenv("CONCOURSE_NETWORK_NAME")
	deployment.CloudConfig, _ = strconv.ParseBool(os.Getenv("BOSH_CLOUD_CONFIG"))
	instances, _ := strconv.Atoi(os.Getenv("CONCOURSE_WEB_INSTANCES"))
	deployment.WebInstances = instances
	webIPs := os.Getenv("CONCOURSE_WEB_IPS")
	if webIPs != "" {
		deployment.WebIPs = strings.Split(webIPs, ",")
	}
	deployment.NetworkRange = os.Getenv("CONCOURSE_NETWORK_RANGE")
	deployment.NetworkGateway = os.Getenv("CONCOURSE_NETWORK_GATEWAY")
	webAzs := os.Getenv("CONCOURSE_WEB_AZS")
	if webAzs != "" {
		deployment.WebAZs = strings.Split(webAzs, ",")
	}
	dbAzs := os.Getenv("CONCOURSE_DATABASE_AZS")
	if dbAzs != "" {
		deployment.DatabaseAZs = strings.Split(dbAzs, ",")
	}
	workerAzs := os.Getenv("CONCOURSE_WORKER_AZS")
	if workerAzs != "" {
		deployment.WorkerAZs = strings.Split(workerAzs, ",")
	}

	var yamlString string
	if err := deployment.Initialize(); err == nil {
		yamlString, err = enaml.Paint(deployment)
		fmt.Println(yamlString)
	} else {
		panic(err.Error())
	}
}
