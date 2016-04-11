package releases

import "github.com/xchapter7x/enaml"

func NewBoshVSphereCPI(version, sha string) enaml.Release {
	return enaml.Release{
		Name: "bosh-vsphere-cpi",
		URL:  "https://bosh.io/d/github.com/cloudfoundry-incubator/bosh-vsphere-cpi-release?v=" + version,
		SHA1: sha,
	}
}
