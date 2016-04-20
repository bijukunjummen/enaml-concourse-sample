package main

import (
	"os"
	"strconv"

	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/deployment/concourse"
)

func main() {
	var instances, _ = strconv.Atoi(os.Getenv("WEB_INSTANCES"))
	var url = os.Getenv("CONCOURSE_URL")
	var username = os.Getenv("CONCOURSE_USERNAME")
	var password = os.Getenv("CONCOURSE_PASSWORD")
	enaml.Paint(concourse.NewDeployment(instances, url, username, password))
}
