# Sample usage of enaml for deployment manifest creation and maintenance

### please see github.com/xchapter7x/enaml

[![wercker status](https://app.wercker.com/status/2202e209e7cc396a6c3e76a450dfcc8b/s/master "wercker status")](https://app.wercker.com/project/bykey/2202e209e7cc396a6c3e76a450dfcc8b)

yes we can actually build, test, compile, etc our deployments and cloud-configs :)

### this is just an example, not guaranteed to create a usable deployment
manifest

```
$ go run main.go > concourse-deployment.yml

or

# we now have a binary release of our cross iaas deployment manifest
$ go build main.go -o concourse-deployment
$ ./concourse-deployment > concourse-deployment.yml
```




### basically your deployment now looks like this
```golang

package main

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/releasejobs"
)

var (
	DefaultName   = "concourse"
	DirectorUUID  = "asdfasdfasdf"
	StemcellAlias = "trusty"
)

func main() {
	enaml.Paint(NewDeployment())
}

type Deployment struct {
	enaml.Deployment
	Manifest *enaml.DeploymentManifest
}

func NewDeployment() (d Deployment) {
	d = Deployment{}
	d.Manifest = new(enaml.DeploymentManifest)
	d.Manifest.SetName(DefaultName)
	d.Manifest.SetDirectorUUID(DirectorUUID)
	d.Manifest.AddReleaseByName("concourse")
	d.Manifest.AddReleaseByName("garden-linux")
	d.Manifest.AddStemcellByName("ubuntu-trusty", StemcellAlias)
	web := enaml.NewInstanceGroup("web", 1, "web", StemcellAlias)
	web.AddAZ("z1")
	web.AddNetwork(enaml.InstanceGroupNetwork{"name": "private"})
	atc := enaml.NewInstanceJob("atc", "concourse", releasejobs.Atc{
		ExternalUrl:        "something",
		BasicAuthUsername:  "user",
		BasicAuthPassword:  "password",
		PostgresqlDatabase: "&atc_db atc",
	})
	tsa := enaml.NewInstanceJob("tsa", "concourse", releasejobs.Tsa{})
	web.AddJob(atc)
	web.AddJob(tsa)
	db := enaml.NewInstanceGroup("db", 1, "database", StemcellAlias)
	worker := enaml.NewInstanceGroup("worker", 1, "worker", StemcellAlias)
	d.Manifest.AddInstanceGroup(web)
	d.Manifest.AddInstanceGroup(db)
	d.Manifest.AddInstanceGroup(worker)
	return
}

func (s Deployment) GetDeployment() enaml.DeploymentManifest {
	return *s.Manifest
}
```
