package stemcells

import "github.com/xchapter7x/enaml"

var NewUbuntuTrusty = func(version, sha1 string) enaml.Stemcell {
	return enaml.Stemcell{
		URL:  "https://bosh.io/d/stemcells/bosh-vsphere-esxi-ubuntu-trusty-go_agent?v=" + version,
		SHA1: sha1,
	}
}
