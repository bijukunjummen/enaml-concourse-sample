package releases

import "github.com/xchapter7x/enaml"

func NewGarden(version, sha string) enaml.Release {
	return enaml.Release{
		Name: "garden-linux",
		URL:  "https://bosh.io/d/github.com/cloudfoundry-incubator/garden-linux-release?v=" + version,
		SHA1: sha,
	}
}
