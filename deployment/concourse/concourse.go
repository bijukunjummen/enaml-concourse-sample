package concourse

import (
	"github.com/xchapter7x/enaml"
	"github.com/xchapter7x/enaml-concourse-sample/standard-components/diskpools"
	"github.com/xchapter7x/enaml-concourse-sample/standard-components/networks"
	"github.com/xchapter7x/enaml-concourse-sample/standard-components/releases"
	"github.com/xchapter7x/enaml-concourse-sample/standard-components/resourcepools"
	"github.com/xchapter7x/enaml-concourse-sample/standard-components/stemcells"
)

var (
	DefaultName            = "concourse"
	ConcourseVersion       = "1.0.0"
	ConcourseSHA           = "xyz"
	GardenVersion          = "1.0.0"
	GardenSHA              = "xyz"
	VSphereCPIVersion      = "1.0.0"
	VSphereCPISHA          = "xyz"
	VSphereStemcellVersion = "1.0.0"
	VSphereStemcellSHA     = "xyz"
)

type Deployment struct {
	enaml.Deployment
	enaml.DeploymentManifest
}

func NewDeployment() (deployment Deployment) {
	deployment = Deployment{}
	deployment.DeploymentManifest = enaml.DeploymentManifest{
		Name: DefaultName,
	}
	deployment.DeploymentManifest.Releases = []enaml.Release{
		releases.NewConcourse(ConcourseVersion, ConcourseSHA),
		releases.NewGarden(GardenVersion, GardenSHA),
	}
	deployment.DeploymentManifest.Networks = []enaml.DeploymentNetwork{
		networks.NewFooNetworkExternal(DefaultName),
	}
	deployment.DeploymentManifest.ResourcePools = []enaml.ResourcePool{
		resourcepools.NewSmallResource(DefaultName, DefaultName),
	}
	deployment.DiskPools = []enaml.DiskPool{
		diskpools.NewDiskPool("db", 10240),
	}
	return
}

func (s Deployment) VSphere() enaml.DeploymentManifest {
	s.Releases = append(s.Releases, releases.NewBoshVSphereCPI(VSphereCPIVersion, VSphereCPISHA))
	for i := range s.ResourcePools {
		s.ResourcePools[i].Stemcell = stemcells.NewUbuntuTrusty(VSphereStemcellVersion, VSphereStemcellSHA)
	}
	return s.DeploymentManifest
}

func (s Deployment) AWS() enaml.DeploymentManifest {
	panic("un-implemented iaas")
}

func (s Deployment) Azure() enaml.DeploymentManifest {
	panic("un-implemented iaas")
}

func (s Deployment) OpenStack() enaml.DeploymentManifest {
	panic("un-implemented iaas")
}
