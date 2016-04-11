package releases

import "github.com/xchapter7x/enaml"

func NewConcourse(version, sha string) enaml.Release {
	return enaml.Release{
		Name: "concourse",
		URL:  "https://bosh.io/d/github.com/concourse/concourse?v=" + version,
		SHA1: sha,
	}
}
