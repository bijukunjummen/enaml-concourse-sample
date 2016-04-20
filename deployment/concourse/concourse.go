package concourse

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/releasejobs"
)

var (
	DefaultName   = "concourse"
	DirectorUUID  = "asdfasdfasdf"
	StemcellAlias = "trusty"
)

type Deployment struct {
	enaml.Deployment
	Manifest *enaml.DeploymentManifest
}

func NewDeployment(webInstances int, url, username, pass string) (d Deployment) {
	if !isStrongPass(pass) {
		panic("Sorry. The given password is too weak")
	}
	d = Deployment{}
	d.Manifest = new(enaml.DeploymentManifest)
	d.Manifest.SetName(DefaultName)
	d.Manifest.SetDirectorUUID(DirectorUUID)
	d.Manifest.AddReleaseByName("concourse")
	d.Manifest.AddReleaseByName("garden-linux")
	d.Manifest.AddStemcellByName("ubuntu-trusty", StemcellAlias)
	web := enaml.NewInstanceGroup("web", insureHAInstanceCount(webInstances), "web", StemcellAlias)
	web.AddAZ("z1")
	web.AddNetwork(enaml.InstanceGroupNetwork{"name": "private"})
	atc := enaml.NewInstanceJob("atc", "concourse", releasejobs.Atc{
		ExternalUrl:        url,
		BasicAuthUsername:  username,
		BasicAuthPassword:  pass,
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

func isStrongPass(pass string) (ok bool) {
	ok = false
	if len(pass) > 10 {
		ok = true
	}
	return
}

func insureHAInstanceCount(instances int) int {
	if instances < 2 {
		instances = 2
	}
	return instances
}

func (s Deployment) GetDeployment() enaml.DeploymentManifest {
	return *s.Manifest
}
