package concourse

import "github.com/xchapter7x/enaml"

var (
	DefaultName  = "concourse"
	DirectorUUID = "asdfasdfasdf"
)

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
	d.Manifest.AddStemcellByName("ubuntu-trusty", "trusty")
	return
}

func (s Deployment) VSphere() enaml.DeploymentManifest {
	return *s.Manifest
}
